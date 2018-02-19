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

func KeyLengthHammingDistance(b []byte, l int) int {
	b1 := b[0:l]
	b2 := b[l : 2*l]
	return HammingDistance(b1, b2)
}

func ProbableKeyLengths(b []byte) []int {
	lim := len(b) / 2
	if lim > 40 {
		lim = 40
	}

	type lengthPair struct {
		l int
		s float64
	}

	lengths := []lengthPair{}
	for i := 2; i <= lim; i++ {
		lengths = append(lengths, lengthPair{i, float64(KeyLengthHammingDistance(b, i)) / float64(i)})
	}

	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i].s < lengths[j].s
	})

	ret := []int{}
	for _, l := range lengths {
		ret = append(ret, l.l)
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
