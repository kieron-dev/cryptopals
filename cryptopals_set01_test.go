package cryptopals_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/operations"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crytopals Set 01", func() {
	It("question 1", func() {
		bytes, err := conversion.HexToBytes("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
		Expect(err).NotTo(HaveOccurred())
		Expect(conversion.BytesToBase64(bytes)).
			To(Equal("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"))

		fmt.Println(string(bytes))
	})

	It("question 2", func() {
		hex1 := "1c0111001f010100061a024b53535009181c"
		hex2 := "686974207468652062756c6c277320657965"
		expectedXor := "746865206b696420646f6e277420706c6179"

		bytes1, err := conversion.HexToBytes(hex1)
		Expect(err).NotTo(HaveOccurred())
		bytes2, err := conversion.HexToBytes(hex2)
		Expect(err).NotTo(HaveOccurred())

		xoredBytes := operations.Xor(bytes1, bytes2)
		Expect(conversion.BytesToHex(xoredBytes)).To(Equal(expectedXor))

		fmt.Println(string(bytes2))
		fmt.Println(string(xoredBytes))
	})

	It("question 3", func() {
		hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
		bytes, err := conversion.HexToBytes(hex)
		Expect(err).NotTo(HaveOccurred())
		clear := operations.SingleCharXorDecrypt(bytes)
		Expect(clear).ToNot(BeEmpty())
		fmt.Println(clear)
	})
})
