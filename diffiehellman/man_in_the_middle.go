package diffiehellman

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

type ManInTheMiddle struct {
	clientComms *DHPeer
	serverComms *DHPeer
	server      *DHPeer
	SentMsg     []byte
	ReturnedMsg []byte
}

func NewManInTheMiddle(server *DHPeer) *ManInTheMiddle {
	m := new(ManInTheMiddle)
	m.server = server
	return m
}

// client -> mitm
func (m *ManInTheMiddle) makeChannel() chan string {
	m.clientComms = NewServer()
	return m.clientComms.makeChannel()
}

// mitm -> server
func (m *ManInTheMiddle) Connect(server DHUser) (clientErr, serverErr error) {
	ch := m.server.makeChannel()
	m.serverComms.ch = ch

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, clientErr = m.initKeyExchange()
	}()

	go func() {
		defer wg.Done()
		_, serverErr = server.completeKeyExchange()
	}()

	wg.Wait()
	return clientErr, serverErr
}

// mitm -> server
func (m *ManInTheMiddle) initKeyExchange() (session []byte, err error) {
	if m.serverComms.ch == nil {
		return nil, errors.New("You need to get the channel first")
	}

	m.serverComms.ch <- fmt.Sprintf("%x", m.clientComms.p)
	m.serverComms.ch <- fmt.Sprintf("%x", m.clientComms.g)
	// hack here...
	m.serverComms.ch <- fmt.Sprintf("%x", m.clientComms.p)

	otherKeyStr := <-m.serverComms.ch
	otherKey := new(big.Int)
	otherKey.SetString(otherKeyStr, 16)

	m.serverComms.sess = m.serverComms.k.Session(otherKey).Bytes()
	return m.serverComms.sess, nil
}

// client -> mitm
func (m *ManInTheMiddle) completeKeyExchange() (session []byte, err error) {
	if m.clientComms.ch == nil {
		return nil, errors.New("you need to get a channel first")
	}

	m.clientComms.g = new(big.Int)
	m.clientComms.p = new(big.Int)
	otherKey := new(big.Int)

	pStr := <-m.clientComms.ch
	gStr := <-m.clientComms.ch
	otherKeyStr := <-m.clientComms.ch

	m.clientComms.p.SetString(pStr, 16)
	m.clientComms.g.SetString(gStr, 16)
	otherKey.SetString(otherKeyStr, 16)

	m.serverComms = NewClient(m.clientComms.p, m.clientComms.g)
	m.Connect(m.server)

	m.clientComms.k = New(m.clientComms.g, m.clientComms.p)

	m.clientComms.ch <- fmt.Sprintf("%x", m.clientComms.p)
	return nil, nil
}

func (m *ManInTheMiddle) ReplyTestMessage() error {
	otherIVStr := <-m.clientComms.ch
	encMsgStr := <-m.clientComms.ch

	sentMsg, err := decodeMsg(encMsgStr, otherIVStr)
	if err != nil {
		return err
	}
	m.SentMsg = sentMsg

	go m.server.ReplyTestMessage()

	m.serverComms.ch <- otherIVStr
	m.serverComms.ch <- encMsgStr

	respIVStr := <-m.serverComms.ch
	respMsgStr := <-m.serverComms.ch

	returnedMsg, err := decodeMsg(respMsgStr, respIVStr)
	if err != nil {
		return err
	}
	m.ReturnedMsg = returnedMsg

	m.clientComms.ch <- respIVStr
	m.clientComms.ch <- respMsgStr

	return nil
}

func decodeMsg(encMsgStr, otherIVStr string) ([]byte, error) {
	blocksize := 16
	sum := sha1.Sum([]byte{})
	key := sum[:16]
	otherIV, err := conversion.HexToBytes(otherIVStr)
	if err != nil {
		return nil, err
	}
	encMsg, err := conversion.HexToBytes(encMsgStr)
	if err != nil {
		return nil, err
	}
	clear, err := operations.AES128CBCDecode(encMsg, key, otherIV)
	if err != nil {
		return nil, err
	}
	return operations.RemovePKCS7(clear, blocksize), nil
}
