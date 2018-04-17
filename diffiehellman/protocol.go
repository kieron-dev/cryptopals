package diffiehellman

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"math/big"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

func StartSession(p, g *big.Int, ch chan string) []byte {
	k := New(g, p)
	ch <- fmt.Sprintf("%x", p)
	ch <- fmt.Sprintf("%x", g)
	ch <- fmt.Sprintf("%x", k.Key)

	otherKeyStr := <-ch
	otherKey := new(big.Int)
	otherKey.SetString(otherKeyStr, 16)

	return k.Session(otherKey).Bytes()
}

func PairSession(ch chan string) []byte {
	g := new(big.Int)
	p := new(big.Int)
	otherKey := new(big.Int)

	pStr := <-ch
	gStr := <-ch
	otherKeyStr := <-ch

	p.SetString(pStr, 16)
	g.SetString(gStr, 16)
	otherKey.SetString(otherKeyStr, 16)

	k := New(g, p)

	ch <- fmt.Sprintf("%x", k.Key)
	return k.Session(otherKey).Bytes()
}

func SendTestMessage(ch chan string, sess []byte, msg []byte) (ok bool, err error) {
	const blocksize = 16
	iv := operations.RandomSlice(blocksize)
	sum := sha1.Sum(sess)
	key := sum[:blocksize]
	paddedMsg := operations.PKCS7(msg, blocksize)
	encrypted, err := operations.AES128CBCEncode(paddedMsg, key, iv)
	if err != nil {
		return false, err
	}
	ch <- fmt.Sprintf("%x", iv)
	ch <- fmt.Sprintf("%x", encrypted)

	respIVStr := <-ch
	respMsgStr := <-ch

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

func ReplyTestMessage(ch chan string, sess []byte) error {
	iv := operations.RandomSlice(16)
	sum := sha1.Sum(sess)
	key := sum[:16]

	otherIVStr := <-ch
	encMsgStr := <-ch

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

	ch <- fmt.Sprintf("%x", iv)
	ch <- fmt.Sprintf("%x", enc)

	return nil
}
