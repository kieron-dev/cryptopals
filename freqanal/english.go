package freqanal

import (
	"sort"
	"strings"
)

var runeOrders map[rune]int

func init() {
	runeOrders = map[rune]int{}

	for i, r := range "etaoinshrdlcumwfgypbvkjxqz" {
		runeOrders[r] = i
	}
}

func FreqScoreEnglish(in string) int {
	score := 0
	sorted := FreqSortStr(in)

	for i, r := range sorted {
		diff := runeOrders[r] - i
		if diff < 0 {
			diff *= -1
		}
		switch diff {
		case 0:
			score += 20
		case 1:
			score += 10
		case 2:
			score += 5
		case 3:
			score += 2
		case 4:
			score += 1
		}
	}

	return score
}

func FreqSortStr(s string) string {
	freqs := GetFreqs(s)
	runes := []rune{}
	for k := range freqs {
		runes = append(runes, k)
	}

	sort.Slice(runes, func(i, j int) bool {
		if freqs[runes[i]] == freqs[runes[j]] {
			return runes[i] > runes[j]
		}
		return freqs[runes[i]] > freqs[runes[j]]
	})
	return string(runes)
}

func GetFreqs(s string) map[rune]int {
	s = strings.ToLower(s)
	ret := map[rune]int{}
	for _, r := range s {
		if r < 'a' || r > 'z' {
			continue
		}
		ret[r]++
	}
	return ret
}
