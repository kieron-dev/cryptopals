package operations

import (
	"strings"

	"github.com/kieron-pivotal/cryptopals/freqanal"
)

func SingleCharXorDecrypt(bytes []byte) string {
	maxScore := 0
	ret := ""

	body := func(r rune) {
		xorString := strings.Repeat(string(r), len(bytes))
		xored := Xor(bytes, []byte(xorString))
		score := freqanal.FreqScoreEnglish(string(xored))
		if score > maxScore {
			maxScore = score
			ret = string(xored)
		}
	}

	for r := 'A'; r <= 'Z'; r++ {
		body(r)
	}
	for r := 'a'; r <= 'z'; r++ {
		body(r)
	}

	return ret
}
