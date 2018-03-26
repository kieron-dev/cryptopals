package md4_test

import (
	"github.com/kieron-pivotal/cryptopals/md4"
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("MD4 MAC", func() {
	var (
		key      []byte
		contents []byte
		mac1     string
	)

	BeforeEach(func() {
		key = operations.RandomSlice(16)
		contents = operations.RandomSlice(1000)
		mac1 = md4.GenerateMD4MAC(key, contents)
	})

	Context("generating MAC", func() {

		It("produces the same MAC for same key and contents", func() {
			mac := md4.GenerateMD4MAC(key, contents)
			Expect(mac).To(Equal(mac1))
		})

		It("varying the contents changes the MAC", func() {
			contents[0], contents[1] = contents[1], contents[0]
			mac := md4.GenerateMD4MAC(key, contents)
			Expect(mac).ToNot(Equal(mac1))
		})

		It("extending the contents changes the MAC", func() {
			contents = append(contents, operations.RandomSlice(100)...)
			mac := md4.GenerateMD4MAC(key, contents)
			Expect(mac).ToNot(Equal(mac1))
		})

		It("varying the key changes the MAC", func() {
			key[0], key[1] = key[1], key[0]
			mac := md4.GenerateMD4MAC(key, contents)
			Expect(mac).ToNot(Equal(mac1))
		})

	})

	Context("verifying MAC", func() {

		It("can verify a MAC", func() {
			ok := md4.VerifyMD4MAC(mac1, key, contents)
			Expect(ok).To(BeTrue(), "Verified MAC")
		})

	})

	Context("reproducing padding", func() {
		DescribeTable("sums a manually padded content same as auto-padded", func(in string) {
			content := []byte(in)
			padding := md4.GetMD4Padding(len(content))
			Expect((len(content) + len(padding)) % 64).To(Equal(0))

			sumManual := md4.SumWithoutPadding(append(content, padding...))
			sumAuto := md4.Sum(content)

			Expect(sumManual).To(Equal(sumAuto))
		},
			Entry("<empty>", ""),
			Entry("foo bar sha", "foo bar sha"),
			Entry("The rain in Spain falls mainly on the plain", "The rain in Spain falls mainly on the plain"),
		)
	})

	Context("extension hack", func() {

		It("can seed using output", func() {
			out := []byte{100, 127, 74, 244, 229, 11, 215, 160, 69, 219, 225, 152, 193, 51, 109, 215}
			in := []uint32{4098522980, 2698447845, 2564938565, 3614258113}

			split := md4.SplitSum(out)
			Expect(split).To(Equal(in))
		})

		It("can extend a md4 prefix key MAC with known key length", func() {
			keyLen := 23
			key := operations.RandomSlice(keyLen)
			orig := []byte("Oh, I do like to be beside the seaside!")
			extension := []byte(" Oh, I do like to be beside the sea!")

			sum := md4.GenerateMD4MAC(key, orig)
			Expect(md4.VerifyMD4MAC(sum, key, orig)).To(BeTrue(), "normal MAC")

			padding := md4.GetMD4Padding(len(orig) + keyLen)
			paddedOrig := append(orig, padding...)
			newSum := md4.ExtendSum(extension, sum, uint64(keyLen+len(orig)+len(padding)))

			newContent := append(paddedOrig, extension...)

			manualSum := md4.GenerateMD4MAC(key, newContent)
			Expect(newSum).To(Equal(manualSum))

			Expect(md4.VerifyMD4MAC(newSum, key, newContent)).To(BeTrue(), "extension MAC")
		})
	})

	It("can get the key length", func() {
		keyLen := 23
		key := operations.RandomSlice(keyLen)
		orig := []byte("Oh, I do like to be beside the seaside!")

		sum := md4.GenerateMD4MAC(key, orig)
		calcKeyLen, err := md4.GetKeyLen(string(orig), sum, func(clear, hash string) bool {
			return md4.VerifyMD4MAC(hash, key, []byte(clear))
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(calcKeyLen).To(Equal(23))
	})
})
