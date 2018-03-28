package sha1

import (
	"bytes"
	"crypto/sha1"
	"io/ioutil"
	"os"

	"github.com/kieron-pivotal/cryptopals/operations"
)

const blocksize = 64

var (
	opad = bytes.Repeat([]byte{0x5c}, blocksize)
	ipad = bytes.Repeat([]byte{0x36}, blocksize)
)

func HMAC(key, contents []byte) []byte {
	if len(key) > blocksize {
		hashedKey := sha1.Sum(key)
		key = hashedKey[:]
	}
	if len(key) < blocksize {
		key = append(key, bytes.Repeat([]byte{0}, blocksize-len(key))...)
	}

	toHash := operations.Xor(key, opad)

	toHash2 := operations.Xor(key, ipad)
	toHash2 = append(toHash2, contents...)
	hash2 := sha1.Sum(toHash2)

	toHash = append(toHash, hash2[:]...)
	hash := sha1.Sum(toHash)
	return hash[:]
}

func FileHMAC(key []byte, filepath string) (hmac []byte, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return hmac, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return hmac, err
	}

	return HMAC(key, b), nil
}
