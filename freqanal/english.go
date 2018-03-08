package freqanal

var runeOrders map[rune]int
var expectedProportions = []float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015,
	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749,
	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758,
	0.00978, 0.02360, 0.00150, 0.01974, 0.00074,
}

func FreqScoreEnglish(in []byte) float64 {

	freqs := map[int]float64{}

	l := float64(0)
	ignored := 0

	for _, r := range in {
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}

		if r >= 'a' && r <= 'z' {
			freqs[int(r-'a')]++
			l++
		} else if r >= 'A' && r <= 'Z' {
			freqs[int(r-'A')]++
			l++
		} else if r == ',' || r == '.' || r == '\'' || r == ' ' ||
			r == '?' || r == ':' || r == ';' {
			l++
		} else if r >= 32 && r <= 126 {
			ignored++
		} else if r == 9 || r == 10 || r == 13 {
			l++
		} else {
			return 1e20
		}
	}

	score := float64(ignored * ignored * ignored)

	for i := 0; i < 26; i++ {
		exp := expectedProportions[i] * l
		fv := freqs[i]
		score += (fv - exp) * (fv - exp) / exp
	}

	return score
}
