package sha1_test

import (
	"github.com/kieron-pivotal/cryptopals/operations"
	"github.com/kieron-pivotal/cryptopals/sha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MAC", func() {
	var (
		key      []byte
		contents []byte
		mac1     string
	)

	BeforeEach(func() {
		key = operations.RandomSlice(16)
		contents = operations.RandomSlice(1000)
		mac1 = sha1.GenerateSHA1MAC(key, contents)
	})

	Context("generating MAC", func() {

		It("produces the same MAC for same key and contents", func() {
			mac := sha1.GenerateSHA1MAC(key, contents)
			Expect(mac).To(Equal(mac1))
		})

		It("varying the contents changes the MAC", func() {
			contents[0], contents[1] = contents[1], contents[0]
			mac := sha1.GenerateSHA1MAC(key, contents)
			Expect(mac).ToNot(Equal(mac1))
		})

		It("extending the contents changes the MAC", func() {
			contents = append(contents, operations.RandomSlice(100)...)
			mac := sha1.GenerateSHA1MAC(key, contents)
			Expect(mac).ToNot(Equal(mac1))
		})

		It("varying the key changes the MAC", func() {
			key[0], key[1] = key[1], key[0]
			mac := sha1.GenerateSHA1MAC(key, contents)
			Expect(mac).ToNot(Equal(mac1))
		})

	})

	Context("verifying MAC", func() {

		It("can verify a MAC", func() {
			ok := sha1.VerifySHA1MAC(mac1, key, contents)
			Expect(ok).To(BeTrue(), "Verified MAC")
		})

	})

	Context("reproducing padding", func() {

		It("sums a manually padded content same as auto-padded", func() {
			content := []byte("foo bar yellow submarine")
			padding := sha1.GetSHA1Padding(content)
			sumManual := sha1.SumWithoutPadding(append(content, padding...))
			sumAuto := sha1.Sum(content)
			Expect(sumManual).To(Equal(sumAuto))
		})
	})
})
