package ctr

import (
	"crypto/aes"

	"github.com/kieron-pivotal/cryptopals/operations"
)

func Encode(clear, key []byte, stream Counter) []byte {
	ret := []byte{}

	stream.BlockCount = 0
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(clear); i += aes.BlockSize {
		end := i + aes.BlockSize
		if end > len(clear) {
			end = len(clear)
		}
		clearPortion := clear[i:end]
		dst := make([]byte, aes.BlockSize)
		block.Encrypt(dst, stream.Bytes())
		xor := operations.Xor(dst, clearPortion)[:len(clearPortion)]
		ret = append(ret, xor...)

		stream.Inc()
	}

	return ret
}

func Edit(ciphertext, key []byte, offset int, newtext []byte, c Counter) []byte {
	bs := aes.BlockSize
	c.BlockCount = offset / bs

	encStream := []byte{}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	for c.BlockCount*bs < offset+len(newtext) {
		dst := make([]byte, bs)
		block.Encrypt(dst, c.Bytes())
		encStream = append(encStream, dst...)
		c.Inc()
	}

	encStream = encStream[offset%bs:]
	encStream = encStream[:len(newtext)]
	enc := operations.Xor(newtext, encStream)

	ret := make([]byte, len(ciphertext))
	copy(ret, ciphertext)
	for i := 0; i < len(newtext); i++ {
		ret[offset+i] = enc[i]
	}
	return ret
}
