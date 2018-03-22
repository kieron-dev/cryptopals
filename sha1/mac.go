package sha1

import (
	"bytes"
	"fmt"

	"github.com/kieron-pivotal/cryptopals/conversion"
)

func GenerateSHA1MAC(key, contents []byte) string {
	sum := Sum(append(key, contents...))
	return conversion.BytesToHex(sum[:])
}

func VerifySHA1MAC(mac string, key, contents []byte) bool {
	recomputeMAC := GenerateSHA1MAC(key, contents)
	fmt.Printf("mac = %+v\n", mac)
	fmt.Printf("recomputeMAC = %+v\n", recomputeMAC)
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

func ExtendSum(extra []byte, prevSumHex string) string {
	prevSum, err := conversion.HexToBytes(prevSumHex)
	if err != nil {
		panic(err)
	}
	seed := SplitSum(prevSum)
	sum := ExtensionSum(extra, seed)
	return conversion.BytesToHex(sum[:])
}
