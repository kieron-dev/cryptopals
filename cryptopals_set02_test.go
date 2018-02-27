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
		ret := examples.ECBBlockPrependerOracle(examples.ECBBlockPrependerEncode)
		Expect(ret).To(ContainSubstring("I just drove by"))
	})

	It("question 13", func() {
		cookie := examples.GetAdminCookie()
		hash := examples.DecryptCookie(cookie)
		fmt.Println(hash)
		Expect(hash["role"]).To(Equal("admin"))
	})

	It("question 14", func() {
		ret := examples.ECBBlockPrependerWithPrefixOracle(examples.ECBBlockPrependerEncodeWithPrefix)
		Expect(ret).To(ContainSubstring("I just drove by"))
	})

	It("question 15", func() {
		padded := []byte("YELLOW SUBMAR")
		padded = append(padded, []byte{3, 3, 2}...)
		_, err := operations.RemovePKCS7Loudly(padded, 16)
		Expect(err).To(MatchError(ContainSubstring("invalid padding")))
	})

	It("question 16", func() {
		enc := examples.EncodeUserdata(":admin<true:")

		isAdmin := false
		for i := 0; i < len(enc)-11; i++ {
			newEnc := make([]byte, len(enc))
			copy(newEnc, enc)
			newEnc[i] ^= 1
			newEnc[i+6] ^= 1
			newEnc[i+11] ^= 1
			if examples.IsAdmin(newEnc) {
				isAdmin = true
				break
			}
		}
		Expect(isAdmin).To(BeTrue())
	})
})
