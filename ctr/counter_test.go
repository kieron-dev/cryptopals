package ctr_test

import (
	"bytes"

	"github.com/kieron-pivotal/cryptopals/ctr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Counter", func() {

	It("first vals of zero none LE counter are...", func() {
		c := ctr.Counter{
			Nonce: bytes.Repeat([]byte{0}, 8),
		}

		Expect(c.Bytes()).To(Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))

		c.Inc()
		Expect(c.Bytes()).To(Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}))
	})

	It("goes over 256 correctly", func() {
		c := ctr.Counter{
			Nonce: bytes.Repeat([]byte{0}, 8),
		}

		for i := 0; i < 255; i++ {
			c.Inc()
		}
		Expect(c.Bytes()).To(Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0}))

		c.Inc()
		Expect(c.Bytes()).To(Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}))
	})
})
