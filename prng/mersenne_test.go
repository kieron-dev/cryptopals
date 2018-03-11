package prng_test

import (
	"github.com/kieron-pivotal/cryptopals/prng"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mersenne", func() {

	var seed uint32
	var m *prng.Mersenne

	JustBeforeEach(func() {
		m = prng.New(seed)
	})

	It("can provide a new mersenne struct", func() {
		Expect(m).To(BeAssignableToTypeOf(&prng.Mersenne{}))
	})

	It("can get the next 3 numbers", func() {
		Expect(m.Next()).To(Equal(uint32(1927864384)))
		Expect(m.Next()).To(Equal(uint32(1064801546)))
		Expect(m.Next()).To(Equal(uint32(2639064362)))
	})

	It("can get the 3n-th next number", func() {
		n := uint32(624)
		for i := uint32(0); i < 3*n; i++ {
			m.Next()
		}
		m.Next()
		Expect(true).To(BeTrue(), "We didn't panic")
	})
})
