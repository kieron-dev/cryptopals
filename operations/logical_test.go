package operations_test

import (
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logical", func() {

	DescribeTable("Xor",
		func(b1, b2, expectedXor []byte) {
			xor := operations.Xor(b1, b2)
			Expect(xor).To(Equal(expectedXor))
		},

		Entry("empty", []byte{}, []byte{}, []byte{}),
		Entry("single bytes", []byte{0}, []byte{1}, []byte{1}),
		Entry("b0 longer than b1", []byte{0, 1}, []byte{1}, []byte{1, 1}),
		Entry("b1 longer than b0", []byte{1}, []byte{1, 0}, []byte{0, 0}),
	)

})
