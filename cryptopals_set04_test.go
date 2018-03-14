package cryptopals_test

import (
	"bytes"
	"crypto/aes"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/ctr"
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CryptopalsSet04", func() {
	It("question 25", func() {
		enc, err := conversion.ReadBase64File("./assets/04_25.txt")
		Expect(err).NotTo(HaveOccurred())

		key := []byte("YELLOW SUBMARINE")
		clear, err := operations.AES128ECBDecode(enc, key)
		Expect(err).NotTo(HaveOccurred())

		key = operations.RandomSlice(16)
		nonce := operations.RandomSlice(8)
		c := ctr.Counter{Nonce: nonce}

		enc = ctr.Encode(clear, key, c)

		newtext := bytes.Repeat([]byte{0}, len(enc))
		stream := ctr.Edit(enc, key, 0, newtext, c)

		newclear := operations.Xor(stream, enc)
		Expect(newclear).To(Equal(clear))
	})

	It("question 26 - can bit fiddle in CTR", func() {
		enc := ctr.ExampleToken(" admin true ")
		pos, err := ctr.GetVarInputPos()
		Expect(err).NotTo(HaveOccurred())
		enc[pos] ^= ' ' ^ ';'
		enc[pos+6] ^= ' ' ^ '='
		enc[pos+11] ^= ' ' ^ ';'
		Expect(ctr.CheckForAdmin(enc)).To(BeTrue())
	})

	It("question 27 - CBC key and iv the same", func() {
		key := operations.RandomSlice(16)
		clear := "Some day. Some day. Some day. Dominion. Come a time. Some say prayers. I'll say mine"
		enc, err := operations.AES128CBCEncode([]byte(clear), key, key)
		Expect(err).NotTo(HaveOccurred())

		for i := 0; i < 16; i++ {
			enc[aes.BlockSize+i] = 0
			enc[2*aes.BlockSize+i] = enc[i]
		}

		ok, out := operations.AES128CBCSaneDecode(enc, key, key)
		Expect(ok).To(BeFalse())
		k := operations.Xor(out[:16], out[32:48])

		Expect(k).To(Equal(key))
	})
})
