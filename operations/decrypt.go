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

func AES128ECBDecode(in []byte, key []byte) (clear []byte, err error) {
	block, err := aes.NewCipher(key)
	blockSize := block.BlockSize()
	if err != nil {
		return []byte{}, err
	}

	for i := 0; i*blockSize < len(in); i++ {
		block.Decrypt(in[i*blockSize:(i+1)*blockSize],
			in[i*blockSize:(i+1)*blockSize],
		)
	}
	return in, nil
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
