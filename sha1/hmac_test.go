package sha1_test

import (
	"io/ioutil"
	"os"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/sha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hmac", func() {

	DescribeTable("Known HMAC-SHA1 values", func(key, contents []byte, expected string) {
		hmac := sha1.HMAC(key, contents)
		hmacHex := conversion.BytesToHex(hmac)
		Expect(hmacHex).To(Equal(expected))
	},
		Entry("<empty>", []byte{}, []byte{}, "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d"),
		Entry("quick brown", []byte("key"), []byte("The quick brown fox jumps over the lazy dog"), "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9"),
	)

	Context("file hmacs", func() {
		It("returns err if file doesn't exist", func() {
			_, err := sha1.FileHMAC([]byte("key"), "/does/not/exist")
			Expect(err).To(HaveOccurred())
		})

		It("returns known hmac for quick brown file", func() {
			file := "/tmp/hmac-sha1-test"
			err := ioutil.WriteFile(file, []byte("The quick brown fox jumps over the lazy dog"), 0644)
			Expect(err).NotTo(HaveOccurred())
			defer os.Remove(file)

			hmac, err := sha1.FileHMAC([]byte("key"), file)
			Expect(err).NotTo(HaveOccurred())
			hmacHex := conversion.BytesToHex(hmac)
			Expect(hmacHex).To(Equal("de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9"))
		})
	})
})
