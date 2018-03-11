package prng

const (
	w uint32 = 32
	n uint32 = 624
	m uint32 = 397
	r uint32 = 31

	a uint32 = 0x9908b0df
	u uint32 = 11
	d uint32 = 0xffffffff
	s uint32 = 7
	b uint32 = 0x9d2c5680
	t uint32 = 15
	c uint32 = 0xefc60000
	l uint32 = 18

	f uint32 = 1812433253
)

type Mersenne struct {
	x   []uint32
	pos uint32
}

func New(seed uint32) *Mersenne {
	m := Mersenne{}
	m.x = []uint32{seed}
	for i := uint32(1); i < n; i++ {
		xi := f*(m.x[i-1]^(m.x[i-1]>>(w-2))) + i
		m.x = append(m.x, xi)
	}
	m.twist()
	return &m
}

func (g *Mersenne) Next() uint32 {
	if g.pos == n {
		g.twist()
	}
	ret := g.x[g.pos]
	g.pos++
	return ret
}

func (g *Mersenne) twist() {
	umask := (uint32(1)<<u - 1) << (w - u - 1)
	lmask := uint32(1)<<l - 1

	g.x = append(g.x, make([]uint32, n)...)
	for k := uint32(0); k < w; k++ {
		maskedAdd := (g.x[k] & umask) | (g.x[k+1] & lmask)
		next := g.x[m+k] ^ rightApplyA(maskedAdd)
		g.x[n+k] = temper(next)
	}

	g.x = g.x[n:]
	g.pos = 0
}

func rightApplyA(x uint32) uint32 {
	if x&1 == 1 {
		return x >> 1
	}
	return (x >> 1) ^ a
}

func temper(x uint32) uint32 {
	next := x ^ (x>>u)&d
	next = next ^ (next<<s)&b
	next = next ^ (next<<t)&c
	next = next ^ (next >> l)
	return next
}
