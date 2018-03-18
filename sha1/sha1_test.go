package sha1_test

import (
	"crypto/sha1"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sha1", func() {
	It("can calculate a sha1sum of random bytes", func() {
		sum := sha1.Sum(operations.RandomSlice(1000))
		Expect(sum).To(HaveLen(20))
	})

	It("calcs the correct sha1sum of 'foo'", func() {
		sum := sha1.Sum([]byte("foo"))
		hex := conversion.BytesToHex(sum[:])
		Expect(hex).To(Equal("0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"))
	})
})
