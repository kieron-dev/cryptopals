package sha1_test

import (
	"github.com/kieron-pivotal/cryptopals/operations"
	"github.com/kieron-pivotal/cryptopals/sha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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
		DescribeTable("sums a manually padded content same as auto-padded", func(in string) {
			content := []byte(in)
			padding := sha1.GetSHA1Padding(len(content))
			Expect((len(content) + len(padding)) % 64).To(Equal(0))

			sumManual := sha1.SumWithoutPadding(append(content, padding...))
			sumAuto := sha1.Sum(content)

			Expect(sumManual).To(Equal(sumAuto))
		},
			Entry("<empty>", ""),
			Entry("foo bar sha", "foo bar sha"),
			Entry("The rain in Spain falls mainly on the plain", "The rain in Spain falls mainly on the plain"),
		)
	})

	Context("extension hack", func() {

		DescribeTable("break byte slice in uint32 slice", func(in []byte, expected []uint32) {
			Expect(sha1.SplitSum(in)).To(Equal(expected))
		},
			Entry("0 0 0 1", []byte{0, 0, 0, 1}, []uint32{1}),
			Entry("0 0 1 1", []byte{0, 0, 1, 1}, []uint32{257}),
			Entry("0 1 1 1", []byte{0, 1, 1, 1}, []uint32{65793}),
			Entry("1 1 1 1", []byte{1, 1, 1, 1}, []uint32{16843009}),
			Entry("1 1 1 1, 2 2 2 2", []byte{1, 1, 1, 1, 2, 2, 2, 2}, []uint32{16843009, 33686018}),
		)

		It("extend with init params same as base sum", func() {
			var (
				init0 uint32 = 0x67452301
				init1 uint32 = 0xEFCDAB89
				init2 uint32 = 0x98BADCFE
				init3 uint32 = 0x10325476
				init4 uint32 = 0xC3D2E1F0
			)
			seed := []uint32{init0, init1, init2, init3, init4}

			orig := []byte("Oh, I do like to be beside the seaside!")

			sum := sha1.Sum(orig)
			newSum := sha1.ExtensionSum(orig, seed, 0)

			Expect(newSum).To(Equal(sum))
		})

		It("can extend a SHA1 prefix key MAC with known key length", func() {
			keyLen := 23
			key := operations.RandomSlice(keyLen)
			orig := []byte("Oh, I do like to be beside the seaside!")
			extension := []byte(" Oh, I do like to be beside the sea!")

			sum := sha1.GenerateSHA1MAC(key, orig)
			Expect(sha1.VerifySHA1MAC(sum, key, orig)).To(BeTrue(), "normal MAC")

			padding := sha1.GetSHA1Padding(len(orig) + keyLen)
			paddedOrig := append(orig, padding...)
			newSum := sha1.ExtendSum(extension, sum, uint64(keyLen+len(orig)+len(padding)))

			newContent := append(paddedOrig, extension...)

			manualSum := sha1.GenerateSHA1MAC(key, newContent)
			Expect(newSum).To(Equal(manualSum))

			Expect(sha1.VerifySHA1MAC(newSum, key, newContent)).To(BeTrue(), "extension MAC")
		})
	})

	It("can get the key length", func() {
		keyLen := 23
		key := operations.RandomSlice(keyLen)
		orig := []byte("Oh, I do like to be beside the seaside!")

		sum := sha1.GenerateSHA1MAC(key, orig)
		calcKeyLen, err := sha1.GetKeyLen(string(orig), sum, func(clear, hash string) bool {
			return sha1.VerifySHA1MAC(hash, key, []byte(clear))
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(calcKeyLen).To(Equal(23))
	})
})
