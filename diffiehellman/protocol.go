package diffiehellman

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"math/big"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

type DHUser struct {
	ch   chan string
	p, g *big.Int
	k    *Keys
	sess []byte
}

type Client struct {
	DHUser
}

type Server struct {
	DHUser
}

type MITM struct {
	DHUser
}

func NewClient(p, g *big.Int) *Client {
	c := new(Client)
	c.p = p
	c.g = g
	c.k = New(c.g, c.p)
	return c
}

func NewServer() *Server {
	s := new(Server)
	return s
}

func (c *Client) StartSession(ch chan string) (session []byte) {
	c.ch = ch
	c.ch <- fmt.Sprintf("%x", c.p)
	c.ch <- fmt.Sprintf("%x", c.g)
	c.ch <- fmt.Sprintf("%x", c.k.Key)

	otherKeyStr := <-c.ch
	otherKey := new(big.Int)
	otherKey.SetString(otherKeyStr, 16)

	c.sess = c.k.Session(otherKey).Bytes()
	return c.sess
}

func (s *Server) PairSession(ch chan string) (session []byte) {
	s.ch = ch
	s.g = new(big.Int)
	s.p = new(big.Int)
	otherKey := new(big.Int)

	pStr := <-s.ch
	gStr := <-s.ch
	otherKeyStr := <-s.ch

	s.p.SetString(pStr, 16)
	s.g.SetString(gStr, 16)
	otherKey.SetString(otherKeyStr, 16)

	s.k = New(s.g, s.p)

	s.ch <- fmt.Sprintf("%x", s.k.Key)
	s.sess = s.k.Session(otherKey).Bytes()
	return s.sess
}

func (c *Client) SendTestMessage(msg []byte) (ok bool, err error) {
	const blocksize = 16
	iv := operations.RandomSlice(blocksize)
	sum := sha1.Sum(c.sess)
	key := sum[:blocksize]
	paddedMsg := operations.PKCS7(msg, blocksize)
	encrypted, err := operations.AES128CBCEncode(paddedMsg, key, iv)
	if err != nil {
		return false, err
	}
	c.ch <- fmt.Sprintf("%x", iv)
	c.ch <- fmt.Sprintf("%x", encrypted)

	respIVStr := <-c.ch
	respMsgStr := <-c.ch

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

func (s *Server) ReplyTestMessage() error {
	iv := operations.RandomSlice(16)
	sum := sha1.Sum(s.sess)
	key := sum[:16]

	otherIVStr := <-s.ch
	encMsgStr := <-s.ch

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

	s.ch <- fmt.Sprintf("%x", iv)
	s.ch <- fmt.Sprintf("%x", enc)

	return nil
}
