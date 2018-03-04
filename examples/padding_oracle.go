package examples

import (
	"fmt"
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

func RandomEncodePaddedCBC(key []byte) (enc, iv []byte) {
	rand.Seed(time.Now().UnixNano())
	clear := RandomClearText()
	iv = operations.RandomSlice(16)
	padded := operations.PKCS7(clear, 16)
	enc, err := operations.AES128CBCEncode(padded, key, iv)
	if err != nil {
		panic(err)
	}
	c, err := operations.AES128CBCDecode(enc, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("---")
	fmt.Println(string(padded), string(c))
	fmt.Println("---")
	return enc, iv
}

func IsCorrectlyPadded(enc, key, iv []byte) bool {
	padded, err := operations.AES128CBCDecode(enc, key, iv)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(padded))
	_, err = operations.RemovePKCS7Loudly(padded, 16)
	return err == nil
}

func PaddingOracle(enc, key, iv []byte) []byte {
	if len(enc) != 16 {
		panic("can only do 16 bytes at the moment")
	}

	res := []byte{}

	for t := 0; t < 256; t++ {
		cp := make([]byte, 16)
		copy(cp, iv)
		cp[15] ^= byte(t) ^ 1
		if IsCorrectlyPadded(enc, key, cp) {
			res = append([]byte{byte(t)}, res...)
			fmt.Println(byte(t))
			break
		}
	}

	for t := 0; t < 256; t++ {
		cp := make([]byte, 16)
		copy(cp, iv)
		cp[15] ^= res[len(res)-1] ^ 2
		cp[14] ^= byte(t) ^ 2
		if IsCorrectlyPadded(enc, key, cp) {
			res = append([]byte{byte(t)}, res...)
			fmt.Println(byte(t))
			break
		}
	}

	for t := 0; t < 256; t++ {
		cp := make([]byte, 16)
		copy(cp, iv)
		cp[15] ^= res[len(res)-1] ^ 3
		cp[14] ^= res[len(res)-2] ^ 3
		cp[13] ^= byte(t) ^ 3
		if IsCorrectlyPadded(enc, key, cp) {
			res = append([]byte{byte(t)}, res...)
			fmt.Println(byte(t))
			break
		}
	}
	return res
}
