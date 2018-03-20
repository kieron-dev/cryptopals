package sha1

import (
	"bytes"

	"github.com/kieron-pivotal/cryptopals/conversion"
)

func GenerateSHA1MAC(key, contents []byte) string {
	sum := Sum(append(key, contents...))
	return conversion.BytesToHex(sum[:])
}

func VerifySHA1MAC(mac string, key, contents []byte) bool {
	recomputeMAC := GenerateSHA1MAC(key, contents)
	return mac == recomputeMAC
}

func GetSHA1Padding(in []byte) []byte {
	bitLen := len(in) * 8
	lenWithInitialx80 := bitLen + 8

	mod512 := lenWithInitialx80 % 512
	if mod512 > 448 {
		mod512 -= 512
	}
	zerosLen := 448 - mod512

	out := []byte{0x80}
	out = append(out, bytes.Repeat([]byte{0}, zerosLen/8)...)

	bitLen64 := uint64(bitLen)
	bitLenBytes := []byte{}
	for i := 0; i < 8; i++ {
		bitLenBytes = append([]byte{uint8(bitLen64)}, bitLenBytes...)
		bitLen64 >>= 8
	}
	return append(out, bitLenBytes...)
}
