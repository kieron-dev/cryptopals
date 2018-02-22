package operations

import "bytes"

func PKCS7(in []byte, blockSize int) []byte {
	diff := (blockSize - (len(in) % blockSize)) % blockSize
	tail := bytes.Repeat([]byte{byte(diff)}, diff)
	return append(in, tail...)
}

func RemovePKCS7(in []byte, blockSize int) []byte {
	l := len(in)
	lastByte := in[l-1]
	if int(lastByte) > blockSize {
		return in
	}

	for i := 1; i < int(lastByte); i++ {
		if in[l-1-i] != lastByte {
			return in
		}
	}
	return in[0 : l-int(lastByte)]
}
