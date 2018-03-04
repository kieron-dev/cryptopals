package operations

import (
	"bytes"
	"crypto/aes"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/freqanal"
)

func SingleCharXorDecrypt(in []byte) (clear string, xorByte byte, score float64) {
	score = 1e20

	for b := byte(0); b < byte(127); b++ {
		xorBytes := []byte{b}
		xored := Xor(in, xorBytes)
		sc := freqanal.FreqScoreEnglish(string(xored))
		if sc < score {
			score = sc
			clear = string(xored)
			xorByte = b
		}
	}

	return clear, xorByte, score
}

func RepeatingXorDecrypt(in []byte) (clear, key string) {
	probableKeyLengths := ProbableKeyLengths(in)
	minScore := 1e20
	probKey := []byte{}

	for _, l := range probableKeyLengths {

		engScore := float64(0)
		key := []byte{}
		for _, s := range SliceBytes(in, l) {
			_, x, sc := SingleCharXorDecrypt(s)
			engScore += sc
			key = append(key, x)
		}

		if engScore < minScore {
			minScore = engScore
			probKey = key
		}

	}
	clear = string(Xor(in, probKey))
	return clear, string(probKey)
}

func AES128ECBEncode(clear []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	blockSize := block.BlockSize()

	for i := 0; i*blockSize < len(clear); i++ {
		slice := clear[i*blockSize : (i+1)*blockSize]
		dst := make([]byte, blockSize)
		block.Encrypt(dst, slice)
		ciphertext = append(ciphertext, dst...)
	}
	return ciphertext, nil
}

func AES128ECBDecode(in []byte, key []byte) (clear []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	blockSize := block.BlockSize()

	for i := 0; i*blockSize < len(in); i++ {
		slice := in[i*blockSize : (i+1)*blockSize]
		dst := make([]byte, blockSize)
		block.Decrypt(dst, slice)
		clear = append(clear, dst...)
	}

	return clear, nil
}

func AES128CBCEncode(in []byte, key []byte, iv []byte) (ciphertext []byte, err error) {
	prev := iv
	blockSize := len(key)
	for i := 0; i*blockSize < len(in); i++ {
		slice := in[i*blockSize : (i+1)*blockSize]
		xor := Xor(prev, slice)
		dst, err := AES128ECBEncode(xor, key)
		if err != nil {
			return []byte{}, err
		}
		ciphertext = append(ciphertext, dst...)
		prev = dst
	}
	return ciphertext, nil
}

func AES128CBCDecode(ciphertext []byte, key []byte, iv []byte) (clear []byte, err error) {
	prev := iv
	blockSize := len(key)
	for i := 0; i*blockSize < len(ciphertext); i++ {
		slice := ciphertext[i*blockSize : (i+1)*blockSize]
		dst, err := AES128ECBDecode(slice, key)
		if err != nil {
			return []byte{}, err
		}
		clear = append(clear, Xor(prev, dst)...)
		prev = slice
	}
	return clear, nil
}

func AES128RandomEncode(clear []byte) (ciphertext []byte, err error) {
	rand.Seed(time.Now().UnixNano())
	key := RandomSlice(16)
	iv := RandomSlice(16)
	preLen := 5 + rand.Intn(6)
	postLen := 5 + rand.Intn(6)
	preBytes := RandomSlice(preLen)
	postBytes := RandomSlice(postLen)
	in := append(preBytes, clear...)
	in = append(in, postBytes...)
	mode := rand.Intn(2)
	if mode == 0 {
		return AES128ECBEncode(in, key)
	}
	return AES128CBCEncode(in, key, iv)
}

func EncodingUsesECB(encoder func([]byte) ([]byte, error)) bool {
	ciphertext, err := encoder(bytes.Repeat([]byte{0}, 48))
	if err != nil {
		panic(err)
	}
	return DetectAES128ECB(conversion.BytesToHex(ciphertext))
}

func DetectBlockSize(encoder func([]byte) ([]byte, error)) int {
	str := []byte("a")
	enc, err := encoder(str)
	if err != nil {
		panic(err)
	}
	l := len(enc)
	i := 2
	for {
		str = bytes.Repeat([]byte("a"), i)
		enc, err = encoder(str)
		if err != nil {
			panic(err)
		}
		if l != len(enc) {
			return len(enc) - l
		}
		l = len(enc)
		i++
	}
}

func DetectAES128ECB(hex string) bool {
	blocks := map[string]int{}
	blockSize := 32

	for i := 0; i*blockSize < len(hex); i++ {
		s := hex[i*blockSize : (i+1)*blockSize]
		if _, ok := blocks[s]; ok {
			return true
		}
		blocks[s]++
	}
	return false
}
