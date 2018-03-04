package examples_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/examples"
	"github.com/kieron-pivotal/cryptopals/operations"

	. "github.com/onsi/ginkgo"
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
		enc, iv := examples.RandomEncodePaddedCBC(key)
		Expect(enc).ToNot(BeEmpty())
		Expect(len(iv)).To(Equal(16))
	})

	It("verifies correct padding", func() {
		key := operations.RandomSlice(16)
		enc, iv := examples.RandomEncodePaddedCBC(key)
		Expect(examples.IsCorrectlyPadded(enc, key, iv)).To(BeTrue())
		enc = append(enc, byte(23))
		Expect(examples.IsCorrectlyPadded(enc, key, iv)).To(BeFalse())
	})

	FIt("can decrypt the last byte of a single blocksize enc", func() {
		key := operations.RandomSlice(16)
		iv := operations.RandomSlice(16)
		clear := "YELLOW SUBMARINE"
		enc, err := operations.AES128CBCEncode([]byte(clear), key, iv)
		Expect(err).NotTo(HaveOccurred())

		hacked := examples.PaddingOracle(enc, key, iv)
		fmt.Println(string(hacked))
		Expect(hacked[len(hacked)-1]).To(Equal(byte('E')))
		Expect(hacked[len(hacked)-2]).To(Equal(byte('N')))
	})

	XIt("can decrypt using padding oracle", func() {
		key := []byte("one two three fo")
		iv := bytes.Repeat([]byte{0}, 16)
		enc, err := operations.AES128CBCEncode([]byte("Yellow Submarine"), key, iv)
		Expect(err).NotTo(HaveOccurred())
		Expect(examples.PaddingOracle(enc, key, iv)).To(Equal([]byte("Yellow Submarine")))
	})

})
