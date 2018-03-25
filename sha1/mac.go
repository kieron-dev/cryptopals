package sha1

import (
	"bytes"
	"errors"

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

func GetSHA1Padding(l int) []byte {
	bitLen := l * 8
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

func SplitSum(sum []byte) []uint32 {
	out := []uint32{}
	for {
		n := uint32(0)
		for i := 0; i < 4; i++ {
			n <<= 8
			n |= uint32(sum[i])
		}
		out = append(out, n)
		if len(sum) > 4 {
			sum = sum[4:]
		} else {
			break
		}
	}
	return out
}

func ExtendSum(extra []byte, prevSumHex string, prevLenPlusPadding uint64) string {
	prevSum, err := conversion.HexToBytes(prevSumHex)
	if err != nil {
		panic(err)
	}
	seed := SplitSum(prevSum)
	sum := ExtensionSum(extra, seed, prevLenPlusPadding)
	return conversion.BytesToHex(sum[:])
}

func GetKeyLen(clear, hash string, verifyHash func(clear, hash string) bool) (l int, err error) {
	lim := 1000
	for l = 0; l < lim; l++ {
		extension := "foo"
		padding := GetSHA1Padding(len(clear) + l)
		sum := ExtendSum([]byte(extension), hash, uint64(l+len(clear)+len(padding)))
		newContents := append([]byte(clear), padding...)
		newContents = append(newContents, []byte(extension)...)
		if verifyHash(string(newContents), sum) {
			return
		}
	}
	return 0, errors.New("Could find key len below " + string(lim))
}
