package examples

import (
	"bytes"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

var key []byte

var secret, err = conversion.Base64ToBytes(
	`Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`)

func ECBBlockPrependerEncode(in []byte) ([]byte, error) {
	if key == nil {
		key = operations.RandomSlice(16)
	}

	clear := append(in, secret...)

	ciphertext, err := operations.AES128ECBEncode(clear, key)
	if err != nil {
		return []byte{}, err
	}
	return ciphertext, nil
}

func ECBBlockPrependerOracle(encoder func([]byte) ([]byte, error)) []byte {
	isECB := operations.EncodingUsesECB(encoder)
	if !isECB {
		panic("Not ECB")
	}
	blocksize := operations.DetectBlockSize(encoder)
	emptyEnc, err := encoder([]byte{})
	if err != nil {
		panic(err)
	}
	blocks := len(emptyEnc) / blocksize

	ret := []byte{}

	for i := 0; i < blocks; i++ {
		for pos := 0; pos < blocksize; pos++ {
			prefix := bytes.Repeat([]byte{'a'}, blocksize-1-pos)

			enc1, err := encoder(prefix)
			if err != nil {
				panic(err)
			}

		outer:
			for b := 0; b < 256; b++ {
				testIn := append(prefix, ret...)
				test, err := encoder(append(testIn, byte(b)))
				if err != nil {
					panic(err)
				}
				for j := 0; j < blocksize; j++ {
					if test[i*blocksize+j] != enc1[i*blocksize+j] {
						continue outer
					}
				}
				ret = append(ret, byte(b))
				break
			}
		}
	}
	return ret
}
