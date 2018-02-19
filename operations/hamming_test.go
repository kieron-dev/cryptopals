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

		Expect(operations.HammingDistance(s1, s1)).To(Equal(0))
		Expect(operations.HammingDistance(s1, s2)).To(Equal(37))
	})

})
