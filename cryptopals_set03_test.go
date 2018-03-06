package cryptopals_test

import (
	"bytes"
	"fmt"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/ctr"
	"github.com/kieron-pivotal/cryptopals/examples"
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CryptopalsSet03", func() {

	It("question 17", func() {
		enc, iv := examples.EncodeRandomText()
		clear := examples.PaddingOracle(enc, iv, examples.IsCorrectlyPadded)

		fmt.Println("---")
		fmt.Println(string(operations.RemovePKCS7(clear, 16)))
		fmt.Println("---")

		Expect(clear).To(HaveLen(len(enc)))
		Expect(string(clear)).To(MatchRegexp(`^0000`))
	})

	It("question 18", func() {
		enc64 := "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
		enc, err := conversion.Base64ToBytes(enc64)
		if err != nil {
			panic(err)
		}

		key := []byte("YELLOW SUBMARINE")
		counter := ctr.Counter{
			Nonce: bytes.Repeat([]byte{0}, 8),
		}

		clear := ctr.Encode(enc, key, counter)
		fmt.Println(string(clear))

		Expect(string(clear)).To(ContainSubstring("Ice, Ice, baby"))
	})
})
