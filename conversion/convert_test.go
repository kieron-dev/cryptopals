package conversion_test

import (
	"github.com/kieron-pivotal/cryptopals/conversion"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("conversion", func() {

	DescribeTable("Happy path HexToBytes",
		func(hex string, expected []byte) {
			bytes, err := conversion.HexToBytes(hex)
			Expect(err).NotTo(HaveOccurred())
			Expect(bytes).To(Equal(expected))
		},

		Entry("empty string", "", []byte{}),
		Entry("zero", "00", []byte{0}),
		Entry("one", "01", []byte{1}),
		Entry("letters", "ff", []byte{255}),
		Entry("more than 1 byte", "ff00", []byte{255, 0}),
	)

	DescribeTable("Sad path HexToBytes",
		func(hex string) {
			_, err := conversion.HexToBytes(hex)
			Expect(err).To(HaveOccurred())
		},

		Entry("odd string length", "0"),
		Entry("non hex char", "h"),
	)

	DescribeTable("Happy path BytesToHex",
		func(bytes []byte, expected string) {
			Expect(conversion.BytesToHex(bytes)).To(Equal(expected))
		},

		Entry("empty slice", []byte{}, ""),
		Entry("zero byte", []byte{0}, "00"),
		Entry("with a letter", []byte{254}, "fe"),
		Entry("multi chars", []byte{254, 1, 253}, "fe01fd"),
	)
})
