package ctr

type Counter struct {
	// Little-endian nonce
	Nonce []byte

	// Little-endian block counter
	BlockCount int
}

func (c *Counter) Inc() {
	c.BlockCount++
}

func (c *Counter) Bytes() []byte {
	ret := make([]byte, 16)
	for i := 0; i < 8; i++ {
		ret[i] = c.Nonce[i]
	}
	i := c.BlockCount
	for b := 8; i > 0 && b < 16; b++ {
		ret[b] = byte(i % 256)
		i /= 256
	}
	return ret
}
