package prng_test

import (
	"fmt"

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
})
