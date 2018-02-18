package operations

func Xor(bytes1, bytes2 []byte) []byte {
	ret := []byte{}

	l1 := len(bytes1)
	l2 := len(bytes2)
	l := l1
	if l2 > l {
		l = l2
	}

	for i := 0; i < l; i++ {
		var b0, b1 byte
		if i < l1 {
			b0 = bytes1[i]
		}
		if i < l2 {
			b1 = bytes2[i]
		}
		ret = append(ret, b0^b1)
	}

	return ret
}
