package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/operations"
)

type Keys struct {
	a   *big.Int
	Key *big.Int
	p   *big.Int
	g   *big.Int
}

func New(g, p *big.Int) *Keys {
	rand.Seed(time.Now().UnixNano())
	k := Keys{p: p,
		g:   g,
		a:   &big.Int{},
		Key: &big.Int{},
	}
	k.a.SetBytes(operations.RandomSlice(1024))
	k.Key.Exp(k.g, k.a, k.p)
	return &k
}

func (k *Keys) Session(otherKey *big.Int) *big.Int {
	s := big.NewInt(1)
	return s.Exp(otherKey, k.a, k.p)
}
