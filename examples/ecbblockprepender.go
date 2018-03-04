package examples

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
)

var (
	key         []byte
	randomStart []byte
	secret      []byte
)

func init() {
	rand.Seed(time.Now().UnixNano())
	key = operations.RandomSlice(16)
	randomStart = operations.RandomSlice(rand.Intn(256))
	var err error
	secret, err = conversion.Base64ToBytes(
		`Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`)
	if err != nil {
		panic(err)
	}

}

func ECBBlockPrependerEncode(in []byte) ([]byte, error) {
	clear := append(in, secret...)
	clear = operations.PKCS7(clear, 16)

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

	return ECBOracle(encoder)
}

func ECBBlockPrependerEncodeWithPrefix(in []byte) ([]byte, error) {
	clear := append(randomStart, in...)
	clear = append(clear, secret...)
	clear = operations.PKCS7(clear, 16)

	ciphertext, err := operations.AES128ECBEncode(clear, key)
	if err != nil {
		return []byte{}, err
	}
	return ciphertext, nil
}

func ECBBlockPrependerWithPrefixOracle(encoder func([]byte) ([]byte, error)) []byte {
	isECB := operations.EncodingUsesECB(encoder)
	if !isECB {
		panic("Not ECB")
	}

	return ECBOracle(encoder)
}

func ECBOracle(encoder func([]byte) ([]byte, error)) []byte {

	blocksize := operations.DetectBlockSize(encoder)

	emptyEnc, err := encoder([]byte{})
	if err != nil {
		panic(err)
	}

	blocks := len(emptyEnc) / blocksize
	startBlock, prefixPreload := GetStartPosAndPrefixLength(encoder, blocksize, blocks)
	ret := []byte{}

	i := startBlock
outer:
	for {
		added := false
		var test []byte
		for pos := 0; pos < blocksize; pos++ {
			prefix := bytes.Repeat([]byte{'a'}, blocksize-1-pos+prefixPreload)

			enc1, err := encoder(prefix)
			if err != nil {
				panic(err)
			}

		inner:
			for b := 0; b < 256; b++ {
				testIn := append(prefix, ret...)
				test, err = encoder(append(testIn, byte(b)))
				if err != nil {
					panic(err)
				}
				for j := 0; j < blocksize; j++ {
					if test[i*blocksize+j] != enc1[i*blocksize+j] {
						continue inner
					}
				}
				ret = append(ret, byte(b))
				added = true
				break
			}
			if !added {
				break outer
			}
		}
		i++
	}
	return ret
}

func GetStartPosAndPrefixLength(
	encoder func([]byte) ([]byte, error),
	blocksize,
	blocks int) (startBlock, prefixLen int) {

	ebcDetected := false

	i := blocksize*2 - 1

	for !ebcDetected {
		i++
		aaa := bytes.Repeat([]byte{'a'}, i)
		enc, err := encoder(aaa)
		if err != nil {
			panic(err)
		}
		dict := map[string]bool{}

		for j := 0; j+1 < blocks; j++ {
			s := conversion.BytesToHex(enc[j*blocksize : (1+j)*blocksize])
			if dict[s] {
				ebcDetected = true
				startBlock = j - 1
				break
			}
			dict[s] = true
		}
	}
	if !ebcDetected {
		panic("not EBC")
	}
	prefixLen = i - blocksize*2
	return startBlock, prefixLen
}
