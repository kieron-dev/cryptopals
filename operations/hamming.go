package operations

import (
	"sort"
)

func HammingDistance(b1, b2 []byte) (dist int) {
	for _, b := range Xor(b1, b2) {
		for b > 0 {
			if b&1 == 1 {
				dist++
			}
			b >>= 1
		}
	}
	return dist
}

func KeyLengthHammingDistance(b []byte, l int) float64 {
	const numToCheck = 4

	slices := [][]byte{}
	for i := 0; i < numToCheck; i++ {
		slices = append(slices, b[l*i:l*(i+1)])
	}

	sum := 0
	for i := 0; i < numToCheck; i++ {
		for j := i + 1; j < numToCheck; j++ {
			sum += HammingDistance(slices[i], slices[j])
		}
	}

	combs := (numToCheck * (numToCheck - 1)) / 2
	return float64(sum) / (float64(combs * l))
}

func ProbableKeyLengths(b []byte) []int {
	lim := len(b) / 4
	if lim > 40 {
		lim = 40
	}

	type lengthPair struct {
		l int
		s float64
	}

	lengths := []lengthPair{}
	for i := 2; i <= lim; i++ {
		lengths = append(lengths, lengthPair{i, KeyLengthHammingDistance(b, i)})
	}

	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i].s < lengths[j].s
	})

	ret := []int{}
	const numToReturn = 3
	for i := 0; i < numToReturn; i++ {
		ret = append(ret, lengths[i].l)
	}
	return ret
}

func SliceBytes(b []byte, l int) [][]byte {
	ret := make([][]byte, l)
	for i, v := range b {
		m := i % l
		ret[m] = append(ret[m], v)
	}
	return ret
}
