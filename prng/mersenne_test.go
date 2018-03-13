package prng_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/prng"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mersenne", func() {

	var seed uint32 = 12345
	var m *prng.Mersenne

	JustBeforeEach(func() {
		m = prng.New(seed)
	})

	It("can provide a new mersenne struct", func() {
		Expect(m).To(BeAssignableToTypeOf(&prng.Mersenne{}))
	})

	Context("with default c++ seed", func() {
		BeforeEach(func() {
			seed = 5489
		})

		It("can get the next 3 numbers", func() {
			Expect(m.Next()).To(Equal(uint32(3499211612)))
			Expect(m.Next()).To(Equal(uint32(581869302)))
			Expect(m.Next()).To(Equal(uint32(3890346734)))
		})
	})

	It("can get the 3n-th next number", func() {
		n := uint32(624)
		for i := uint32(0); i < 3*n; i++ {
			m.Next()
		}
		m.Next()
		Expect(true).To(BeTrue(), "We didn't panic")
	})

	It("doesn't give duplicates in 100000 ops", func() {
		d := map[uint32]bool{}
		for i := 0; i < 100000; i++ {
			r := m.Next()
			if d[r] {
				Fail(fmt.Sprintf("Got a dup: %d, on step: %d", r, i))
			}
			d[r] = true
		}
	})

	It("can undo a xor-rshift-and", func() {
		rand.Seed(time.Now().UnixNano())
		in := rand.Uint32()
		s := rand.Intn(8) + 2
		a := rand.Uint32()
		res := prng.XorRshiftAnd(in, uint32(s), a)
		undone := prng.UndoXorRshiftAnd(res, uint32(s), a)
		Expect(undone).To(Equal(in))
	})

	It("can undo a xor-lshift-and", func() {
		rand.Seed(time.Now().UnixNano())
		in := rand.Uint32()
		s := rand.Intn(8) + 2
		a := rand.Uint32()
		res := prng.XorLshiftAnd(in, uint32(s), a)
		undone := prng.UndoXorLshiftAnd(res, uint32(s), a)
		Expect(undone).To(Equal(in))
	})

	It("can detemper a temper", func() {
		in := uint32(1231548362)
		t := prng.Temper(in)
		Expect(prng.Detemper(t)).To(Equal(in))
	})

	It("can be seeded with 264 uint32s", func() {
		mer := prng.Mersenne{}
		seed := []uint32{}
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 624; i++ {
			seed = append(seed, rand.Uint32())
		}
		mer.Seed(seed)
		mer.Next()
		Expect(true).To(BeTrue(), "didn't panic")
	})

	It("can be cloned", func() {
		rand.Seed(time.Now().UnixNano())
		mer1 := prng.New(rand.Uint32())
		var seed []uint32
		for i := 0; i < 624; i++ {
			seed = append(seed, prng.Detemper(mer1.Next()))
		}
		mer2 := prng.Mersenne{}
		mer2.Seed(seed)

		for i := 0; i < 1000; i++ {
			Expect(mer2.Next()).To(Equal(mer1.Next()))
		}
	})

})
