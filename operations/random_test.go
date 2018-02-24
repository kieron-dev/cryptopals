package operations_test

import (
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Random", func() {

	It("can generate a slice of random bytes", func() {
		bytes := operations.RandomSlice(10)
		Expect(bytes).To(HaveLen(10))

		bytes2 := operations.RandomSlice(10)
		Expect(bytes2).ToNot(Equal(bytes))
	})

})
