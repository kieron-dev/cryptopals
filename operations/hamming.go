package operations

func HammingDistance(s1, s2 string) (dist int) {
	for _, b := range Xor([]byte(s1), []byte(s2)) {
		for b > 0 {
			if b&1 == 1 {
				dist++
			}
			b >>= 1
		}
	}
	return dist
}
