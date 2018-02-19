package operations

func Xor(bytes1, bytes2 []byte) []byte {
	ret := []byte{}

	l2 := len(bytes2)

	for i, b0 := range bytes1 {
		b1 := bytes2[i%l2]
		ret = append(ret, b0^b1)
	}

	return ret
}
