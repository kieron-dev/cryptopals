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
