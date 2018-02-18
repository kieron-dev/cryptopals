package operations

import (
	"strings"

	"github.com/kieron-pivotal/cryptopals/freqanal"
)

func SingleCharXorDecrypt(bytes []byte) (string, float64) {
	var minScore float64
	minScore = 100000
	ret := ""

	for r := rune(0); r < rune(127); r++ {
		xorString := strings.Repeat(string(r), len(bytes))
		xored := Xor(bytes, []byte(xorString))
		score := freqanal.FreqScoreEnglish(string(xored))
		if score < minScore {
			minScore = score
			ret = string(xored)
		}
	}

	return ret, minScore
}
