package cryptopals_test

import (
	"bytes"

	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CryptopalsSet02", func() {

	It("question 9", func() {
		in := []byte("YELLOW SUBMARINE")
		out := append(in, bytes.Repeat([]byte{4}, 4)...)
		Expect(operations.PKCS7(in, 20)).To(Equal(out))
	})

})
