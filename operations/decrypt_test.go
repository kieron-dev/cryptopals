package operations_test

import (
	"bytes"

	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decrypt", func() {
	Describe("AES", func() {

		key := []byte("Yellow Submarine")
		clear := []byte("Some Text Really")
		iv := bytes.Repeat([]byte{0}, 16)

		Context("ECB", func() {
			It("can encode and decode something of blocksize length", func() {
				ciphertext, err := operations.AES128ECBEncode(clear, key)
				Expect(err).NotTo(HaveOccurred())
				Expect(ciphertext).ToNot(Equal(clear))

				decrypted, err := operations.AES128ECBDecode(ciphertext, key)
				Expect(err).NotTo(HaveOccurred())
				Expect(decrypted).To(Equal(clear))
			})
		})

		Context("CBC", func() {
			It("can encode and decode something", func() {
				ciphertext, err := operations.AES128CBCEncode(bytes.Repeat(clear, 2), key, iv)
				Expect(err).NotTo(HaveOccurred())
				Expect(ciphertext).ToNot(Equal(clear))

				decoded, err := operations.AES128CBCDecode(ciphertext, key, iv)
				Expect(err).NotTo(HaveOccurred())
				Expect(decoded).To(Equal(bytes.Repeat(clear, 2)))
			})
		})
	})
})
