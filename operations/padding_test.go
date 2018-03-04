package operations_test

import (
	"bytes"
	"fmt"

	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Operations/Padding", func() {

	Context("padding", func() {
		It("pads a reasonable request", func() {
			in := []byte("yellow submarine")
			out := append(in, bytes.Repeat([]byte{4}, 4)...)
			Expect(operations.PKCS7(in, 20)).To(Equal(out))
		})

		It("pads a whole extra block if len % blocksize == 0", func() {
			in := []byte("foo")
			Expect(operations.PKCS7(in, 3)).To(Equal([]byte{'f', 'o', 'o', 3, 3, 3}))
		})

		It("pads a short string", func() {
			in := []byte("foo")
			Expect(operations.PKCS7(in, 4)).To(Equal(append(in, byte(1))))
		})

		It("pads a string longer than blocksize", func() {
			in := []byte("foo bar")
			Expect(operations.PKCS7(in, 4)).To(Equal(append(in, byte(1))))
		})
	})

	Context("removing padding and returning an error on failure", func() {
		It("removes good padding without error", func() {
			padded := []byte{1, 2, 3, 4, 4, 4, 4, 4}
			unpadded, err := operations.RemovePKCS7Loudly(padded, 8)
			Expect(unpadded).To(Equal([]byte{1, 2, 3, 4}))
			Expect(err).ToNot(HaveOccurred())
		})

		It("complains about wrong length", func() {
			wrongLength := []byte{1, 2, 4, 4, 4, 4}
			_, err := operations.RemovePKCS7Loudly(wrongLength, 8)
			Expect(err).To(HaveOccurred())
		})

		It("complains about missing padding", func() {
			missingPadding := []byte("asdfasdf")
			_, err := operations.RemovePKCS7Loudly(missingPadding, 8)
			fmt.Println(err)
			Expect(err).To(HaveOccurred())
		})

		It("complains about a final zero", func() {
			finalZero := []byte{1, 2, 3, 4, 0}
			_, err := operations.RemovePKCS7Loudly(finalZero, 5)
			Expect(err).To(HaveOccurred())
		})

		It("complains about wrong pattern of padding", func() {
			wrongPattern := []byte{1, 2, 3, 4, 3, 4, 3, 4}
			_, err := operations.RemovePKCS7Loudly(wrongPattern, 8)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("quiet padding removal", func() {
		It("removes padding", func() {
			padded := []byte{1, 2, 3, 4, 4, 4, 4, 4}
			Expect(operations.RemovePKCS7(padded, 8)).To(Equal([]byte{1, 2, 3, 4}))
		})

		It("removes full block padding", func() {
			padded := []byte{'a', 's', 'd', 'f', 4, 4, 4, 4}
			Expect(operations.RemovePKCS7(padded, 4)).To(Equal([]byte("asdf")))
		})

		It("returns input when an error occurs", func() {
			padded := []byte("asdf")
			Expect(operations.RemovePKCS7(padded, 4)).To(Equal(padded))
			_, err := operations.RemovePKCS7Loudly(padded, 4)
			Expect(err).To(HaveOccurred())
		})
	})

})
