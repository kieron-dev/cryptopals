package cryptopals_test

import (
	"bytes"
	"fmt"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/examples"
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

	It("question 10", func() {
		ciphertext, err := conversion.ReadBase64File("./assets/02_10.txt")
		Expect(err).NotTo(HaveOccurred())

		iv := bytes.Repeat([]byte{0}, 16)
		key := []byte("YELLOW SUBMARINE")

		clear, err := operations.AES128CBCDecode(ciphertext, key, iv)
		Expect(err).NotTo(HaveOccurred())
		fmt.Println(string(clear))
	})

	It("question 11", func() {
		countCBC := 0
		countECB := 0

		for i := 0; i < 16; i++ {
			if operations.EncodingUsesECB(operations.AES128RandomEncode) {
				countECB++
			} else {
				countCBC++
			}
		}

		Expect(countECB).To(BeNumerically(">", 0))
		Expect(countCBC).To(BeNumerically(">", 0))
	})

	It("question 12", func() {
		cookie := examples.GetAdminCookie()
		hash := examples.DecryptCookie(cookie)
		fmt.Println(hash)
		Expect(hash["role"]).To(Equal("admin"))
	})

})
