package cryptopals_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/examples"
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CryptopalsSet03", func() {

	It("question 17", func() {
		enc, iv := examples.EncodeRandomText()
		clear := examples.PaddingOracle(enc, iv)

		fmt.Println("---")
		fmt.Println(string(operations.RemovePKCS7(clear, 16)))
		fmt.Println("---")

		Expect(clear).To(HaveLen(len(enc)))
		Expect(string(clear)).To(MatchRegexp(`^0000`))
	})

})
