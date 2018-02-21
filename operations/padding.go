package operations

import "bytes"

func PKCS7(in []byte, blockSize int) []byte {
	diff := (blockSize - (len(in) % blockSize)) % blockSize
	tail := bytes.Repeat([]byte{byte(diff)}, diff)
	return append(in, tail...)
}
