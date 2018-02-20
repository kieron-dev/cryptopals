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
	b1 := b[0:l]
	b2 := b[l : 2*l]
	b3 := b[2*l : 3*l]
	b4 := b[3*l : 4*l]
	d12 := float64(HammingDistance(b1, b2))
	d13 := float64(HammingDistance(b1, b3))
	d14 := float64(HammingDistance(b1, b4))
	d23 := float64(HammingDistance(b2, b3))
	d24 := float64(HammingDistance(b2, b4))
	d34 := float64(HammingDistance(b3, b4))
	return (d12 + d13 + d14 + d23 + d24 + d34) / (6 * float64(l))
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
	for _, l := range lengths {
		ret = append(ret, l.l)
	}
	return ret[:3]
}

func SliceBytes(b []byte, l int) [][]byte {
	ret := make([][]byte, l)
	for i, v := range b {
		m := i % l
		ret[m] = append(ret[m], v)
	}
	return ret
}
