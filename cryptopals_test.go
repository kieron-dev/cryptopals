package cryptopals_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/conversion"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crytopals Set 01", func() {
	It("question 1", func() {
		bytes, err := conversion.HexToBytes("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
		fmt.Println(string(bytes))
		Expect(err).NotTo(HaveOccurred())
		Expect(conversion.BytesToBase64(bytes)).
			To(Equal("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"))
	})
})
