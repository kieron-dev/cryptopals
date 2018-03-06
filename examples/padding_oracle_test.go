package examples_test

import (
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/examples"
	"github.com/kieron-pivotal/cryptopals/operations"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("PaddingOracle", func() {

	It("can randomly select a clear text", func() {
		rand.Seed(time.Now().UnixNano())
		cipherText1 := examples.RandomClearText()
		Eventually(examples.RandomClearText).ShouldNot(Equal(cipherText1))
	})

	It("can encode a clear text", func() {
		key := operations.RandomSlice(16)
		enc, iv := examples.EncodePaddedCBC(key)
		Expect(enc).ToNot(BeEmpty())
		Expect(len(iv)).To(Equal(16))
	})

	It("verifies correct padding", func() {
		enc, iv := examples.EncodePaddedCBC([]byte("foo"))
		Expect(examples.IsCorrectlyPadded(enc, iv)).To(BeTrue())
		enc = append(enc, byte(23))
		Expect(examples.IsCorrectlyPadded(enc, iv)).To(BeFalse())
	})

	DescribeTable("can decrypt various strings",
		func(clear string) {
			enc, iv := examples.EncodePaddedCBC([]byte(clear))

			hacked := examples.PaddingOracle(enc, iv)
			unpadded := operations.RemovePKCS7(hacked, 16)
			Expect(string(unpadded)).To(Equal(clear))
		},
		Entry("stuff", "stuff"),
		Entry("Hello, World!", "Hello, World!"),
		Entry("", ""),
		Entry("123456781234567812345678", "123456781234567812345678"),
	)

})
