package operations_test

import (
	"bytes"

	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Operations/Padding", func() {

	It("pads a reasonable request", func() {
		in := []byte("yellow submarine")
		out := append(in, bytes.Repeat([]byte{4}, 4)...)
		Expect(operations.PKCS7(in, 20)).To(Equal(out))
	})

	It("returns the same if already ok", func() {
		in := []byte("foo")
		Expect(operations.PKCS7(in, 3)).To(Equal(in))
	})

	It("pads a short string", func() {
		in := []byte("foo")
		Expect(operations.PKCS7(in, 4)).To(Equal(append(in, byte(1))))
	})

	It("pads a string longer than blocksize", func() {
		in := []byte("foo bar")
		Expect(operations.PKCS7(in, 4)).To(Equal(append(in, byte(1))))
	})

	It("removes padding", func() {
		padded := []byte{1, 2, 3, 4, 4, 4, 4, 4}
		Expect(operations.RemovePKCS7(padded, 8)).To(Equal([]byte{1, 2, 3, 4}))
	})

	It("leaves non-padding", func() {
		padded := []byte{1, 2, 3, 4, 5, 4, 4, 4}
		Expect(operations.RemovePKCS7(padded, 8)).To(Equal(padded))
	})

})
