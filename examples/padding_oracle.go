package examples

import (
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

var encB64s = []string{
	"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
	"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
	"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
	"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
	"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
	"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
	"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
	"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
	"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
	"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
}

var (
	cipherTexts      [][]byte
	paddingOracleKey []byte
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for _, b64 := range encB64s {
		cipherText, err := conversion.Base64ToBytes(b64)
		if err != nil {
			panic(err)
		}
		cipherTexts = append(cipherTexts, cipherText)
	}
	paddingOracleKey = operations.RandomSlice(16)
}

func RandomClearText() []byte {
	return cipherTexts[rand.Intn(len(cipherTexts))]
}

func EncodeRandomText() (enc, iv []byte) {
	return EncodePaddedCBC(RandomClearText())
}

func EncodePaddedCBC(clear []byte) (enc, iv []byte) {
	iv = operations.RandomSlice(16)
	padded := operations.PKCS7(clear, 16)
	enc, err := operations.AES128CBCEncode(padded, paddingOracleKey, iv)
	if err != nil {
		panic(err)
	}
	return enc, iv
}

func IsCorrectlyPadded(enc, iv []byte) bool {
	padded, err := operations.AES128CBCDecode(enc, paddingOracleKey, iv)
	if err != nil {
		panic(err)
	}
	_, err = operations.RemovePKCS7Loudly(padded, 16)
	return err == nil
}

func PaddingOracle(enc, iv []byte) []byte {
	l := len(enc)
	res := make([]byte, l)

	blocksize := 16

	for block := l / blocksize; block > 0; block-- {
		blockToCheck := enc[(block-1)*blocksize : block*blocksize]
		blockToTweak := iv
		if block > 1 {
			blockToTweak = enc[(block-2)*blocksize : (block-1)*blocksize]
		}
		tweakBlock := make([]byte, blocksize)

		for i := 0; i < 16; i++ {
			for j := 0; j < i; j++ {
				tweakBlock[15-j] = blockToTweak[15-j] ^ res[(block-1)*blocksize+15-j] ^ byte(i+1)
			}

			for t := 0; t < 256; t++ {
				tweakBlock[15-i] = blockToTweak[15-i] ^ byte(t) ^ byte(i+1)

				if IsCorrectlyPadded(blockToCheck, tweakBlock) {
					res[(block-1)*blocksize+15-i] = byte(t)
					break
				}
			}
		}
	}
	return res
}
