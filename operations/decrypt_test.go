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

	It("can detect CBC or EBC in a black box", func() {
		key := operations.RandomSlice(16)
		iv := operations.RandomSlice(16)

		Expect(operations.EncodingUsesECB(func(in []byte) ([]byte, error) {
			return operations.AES128ECBEncode(in, key)
		})).To(BeTrue())

		Expect(operations.EncodingUsesECB(func(in []byte) ([]byte, error) {
			return operations.AES128CBCEncode(in, key, iv)
		})).To(BeFalse())
	})

	It("can detect block size", func() {
		encoder := func(in []byte) ([]byte, error) {
			key := operations.RandomSlice(16)
			return operations.AES128ECBEncode(in, key)
		}
		Expect(operations.DetectBlockSize(encoder)).To(Equal(16))
	})
})
