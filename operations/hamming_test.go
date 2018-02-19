package operations_test

import (
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hamming", func() {

	It("calcs hamming distance between two strings", func() {
		s1 := "this is a test"
		s2 := "wokka wokka!!!"

		Expect(operations.HammingDistance([]byte(s1), []byte(s1))).To(Equal(0))
		Expect(operations.HammingDistance([]byte(s1), []byte(s2))).To(Equal(37))
	})

})

var _ = Describe("Hamming in str", func() {

	It("calcs hamming distance in a string", func() {
		Expect(operations.KeyLengthHammingDistance([]byte("this is a testwokka wokka!!!"), 14)).
			To(Equal(37))
	})

})

var _ = Describe("Slice Bytes", func() {
	It("slices bytes", func() {
		bytes := []byte{0, 1, 2, 3, 4, 5, 6, 7}
		slices := operations.SliceBytes(bytes, 3)
		Expect(slices[0]).To(Equal([]byte{0, 3, 6}))
		Expect(slices[1]).To(Equal([]byte{1, 4, 7}))
		Expect(slices[2]).To(Equal([]byte{2, 5}))
	})
})
