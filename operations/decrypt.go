package operations

import (
	"crypto/aes"

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

func AES128ECBEncode(in []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	blockSize := block.BlockSize()

	for i := 0; i*blockSize < len(in); i++ {
		slice := in[i*blockSize : (i+1)*blockSize]
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
