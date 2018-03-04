package operations

import (
	"bytes"
	"errors"
)

func PKCS7(in []byte, blockSize int) []byte {
	diff := (blockSize - (len(in) % blockSize)) % blockSize
	if diff == 0 {
		diff = blockSize
	}
	tail := bytes.Repeat([]byte{byte(diff)}, diff)
	return append(in, tail...)
}

func RemovePKCS7(in []byte, blockSize int) []byte {
	str, _ := RemovePKCS7Loudly(in, blockSize)
	return str
}

func RemovePKCS7Loudly(in []byte, blockSize int) ([]byte, error) {
	l := len(in)
	if l%blockSize != 0 {
		return in, errors.New("length not multiple of blocksize")
	}
	lastByte := in[l-1]

	if int(lastByte) > blockSize {
		return in, errors.New("invalid padding")
	}

	for i := 1; i < int(lastByte); i++ {
		if in[l-1-i] != lastByte {
			return in, errors.New("invalid padding")
		}
	}
	return in[0 : l-int(lastByte)], nil
}
