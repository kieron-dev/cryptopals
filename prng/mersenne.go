package prng

import "fmt"

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

func (g *Mersenne) Seed(seed []uint32) {
	if len(seed) != int(n) {
		panic(fmt.Sprintf("wrong seed length: %d", len(seed)))
	}
	g.x = seed
	g.pos = n
}

func (g *Mersenne) Next() uint32 {
	if g.pos == n {
		g.twist()
	}
	ret := Temper(g.x[g.pos])
	g.pos++
	return ret
}

func (g *Mersenne) twist() {
	umask := (uint32(1)<<(w-r) - 1) << r
	lmask := uint32(1)<<r - 1

	g.x = append(g.x, make([]uint32, n)...)
	for k := uint32(0); k < n; k++ {
		maskedAdd := (g.x[k] & umask) | (g.x[k+1] & lmask)
		g.x[n+k] = g.x[m+k] ^ rightApplyA(maskedAdd)
	}

	g.x = g.x[n:]
	g.pos = 0
}

func rightApplyA(x uint32) uint32 {
	if x&1 == 0 {
		return x >> 1
	}
	return (x >> 1) ^ a
}

func Temper(x uint32) uint32 {
	next := XorRshiftAnd(x, u, d)
	next = XorLshiftAnd(next, s, b)
	next = XorLshiftAnd(next, t, c)
	return XorRshiftAnd(next, l, ^uint32(0))
}

func Detemper(z uint32) uint32 {
	w := UndoXorRshiftAnd(z, l, ^uint32(0))
	w = UndoXorLshiftAnd(w, t, c)
	w = UndoXorLshiftAnd(w, s, b)
	return UndoXorRshiftAnd(w, u, d)
}

func XorRshiftAnd(x, s, a uint32) uint32 {
	return x ^ x>>s&a
}

func XorLshiftAnd(x, s, a uint32) uint32 {
	return x ^ x<<s&a
}

func UndoXorRshiftAnd(r, s, a uint32) uint32 {
	correctBits := s
	w := r
	for correctBits < 32 {
		w >>= s
		w &= a
		w = r ^ w
		correctBits += s
	}
	return w
}

func UndoXorLshiftAnd(r, s, a uint32) uint32 {
	correctBits := s
	w := r
	for correctBits < 32 {
		w <<= s
		w &= a
		w = r ^ w
		correctBits += s
	}
	return w
}
