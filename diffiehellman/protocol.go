package diffiehellman

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

type DHPeer struct {
	ch   chan string
	p, g *big.Int
	k    *Keys
	sess []byte
}

type DHUser interface {
	Connect(p DHUser) (error, error)
	makeChannel() chan string
	initKeyExchange() ([]byte, error)
	completeKeyExchange() ([]byte, error)
}

func (d *DHPeer) Connect(server DHUser) (clientErr, serverErr error) {
	ch := server.makeChannel()
	d.ch = ch

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, clientErr = d.initKeyExchange()
	}()

	go func() {
		defer wg.Done()
		_, serverErr = server.completeKeyExchange()
	}()

	wg.Wait()
	return clientErr, serverErr
}

func (d *DHPeer) makeChannel() chan string {
	ch := make(chan string)
	d.ch = ch
	return ch
}

func NewClient(p, g *big.Int) *DHPeer {
	c := new(DHPeer)
	c.p = p
	c.g = g
	c.k = New(c.g, c.p)
	return c
}

func NewServer() *DHPeer {
	s := new(DHPeer)
	return s
}

func (d *DHPeer) initKeyExchange() (session []byte, err error) {
	if d.ch == nil {
		return nil, errors.New("You need to get the channel first")
	}

	d.ch <- fmt.Sprintf("%x", d.p)
	d.ch <- fmt.Sprintf("%x", d.g)
	d.ch <- fmt.Sprintf("%x", d.k.Key)

	otherKeyStr := <-d.ch
	otherKey := new(big.Int)
	otherKey.SetString(otherKeyStr, 16)

	d.sess = d.k.Session(otherKey).Bytes()
	return d.sess, nil
}

func (d *DHPeer) completeKeyExchange() (session []byte, err error) {
	if d.ch == nil {
		return nil, errors.New("you need to get a channel first")
	}

	d.g = new(big.Int)
	d.p = new(big.Int)
	otherKey := new(big.Int)

	pStr := <-d.ch
	gStr := <-d.ch
	otherKeyStr := <-d.ch

	d.p.SetString(pStr, 16)
	d.g.SetString(gStr, 16)
	otherKey.SetString(otherKeyStr, 16)

	d.k = New(d.g, d.p)

	d.ch <- fmt.Sprintf("%x", d.k.Key)
	d.sess = d.k.Session(otherKey).Bytes()
	return d.sess, nil
}

func (d *DHPeer) SendTestMessage(msg []byte) (ok bool, err error) {
	const blocksize = 16
	iv := operations.RandomSlice(blocksize)
	sum := sha1.Sum(d.sess)
	key := sum[:blocksize]
	paddedMsg := operations.PKCS7(msg, blocksize)
	encrypted, err := operations.AES128CBCEncode(paddedMsg, key, iv)
	if err != nil {
		return false, err
	}
	d.ch <- fmt.Sprintf("%x", iv)
	d.ch <- fmt.Sprintf("%x", encrypted)

	respIVStr := <-d.ch
	respMsgStr := <-d.ch

	iv, err = conversion.HexToBytes(respIVStr)
	if err != nil {
		return false, err
	}
	respMsg, err := conversion.HexToBytes(respMsgStr)
	if err != nil {
		return false, err
	}

	clear, err := operations.AES128CBCDecode(respMsg, key, iv)
	if err != nil {
		return false, err
	}
	clear = operations.RemovePKCS7(clear, blocksize)

	return bytes.Equal(clear, msg), nil
}

func (d *DHPeer) ReplyTestMessage() error {
	iv := operations.RandomSlice(16)
	sum := sha1.Sum(d.sess)
	key := sum[:16]

	otherIVStr := <-d.ch
	encMsgStr := <-d.ch

	otherIV, err := conversion.HexToBytes(otherIVStr)
	if err != nil {
		return err
	}
	encMsg, err := conversion.HexToBytes(encMsgStr)
	if err != nil {
		return err
	}

	clear, err := operations.AES128CBCDecode(encMsg, key, otherIV)
	if err != nil {
		return err
	}

	enc, err := operations.AES128CBCEncode(clear, key, iv)
	if err != nil {
		return err
	}

	d.ch <- fmt.Sprintf("%x", iv)
	d.ch <- fmt.Sprintf("%x", enc)

	return nil
}
