package examples_test

import (
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

	FIt("can decrypt the last byte of a single blocksize enc", func() {
		clear := "Hi, my name is K"
		enc, iv := examples.EncodePaddedCBC([]byte(clear))

		hacked := examples.PaddingOracle(enc, iv)
		fmt.Println(string(hacked))
		Expect(hacked[len(hacked)-1]).To(Equal(byte('K')))
		Expect(hacked[len(hacked)-2]).To(Equal(byte(' ')))
	})

	XIt("can decrypt using padding oracle", func() {
		enc, iv := examples.EncodePaddedCBC([]byte("Yellow Submarine"))
		Expect(examples.PaddingOracle(enc, iv)).To(Equal([]byte("Yellow Submarine")))
	})

})
