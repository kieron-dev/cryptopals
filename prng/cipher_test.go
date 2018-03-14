package prng_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/prng"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cipher", func() {
	It("can encode some text", func() {
		clear := "I wandered lonely as a cloud"
		ciphertext := prng.Encode([]byte(clear), 4321)
		Expect(ciphertext).To(HaveLen(len(clear)))
		Expect([]byte(clear)).ToNot(Equal(ciphertext))
	})

	It("can decrypt given correct key", func() {
		clear := "I wandered lonely as a cloud"
		ciphertext := prng.Encode([]byte(clear), 4321)
		decoded := prng.Decode(ciphertext, 4321)
		Expect(decoded).To(Equal([]byte(clear)))

		wrongDecoded := prng.Decode(ciphertext, 1234)
		Expect(wrongDecoded).ToNot(Equal([]byte(clear)))
	})

	It("can encode clear with a random prefix", func() {
		clear := "I wandered lonely as a cloud"
		ciphertext := prng.EncodeWithRandomPrefix([]byte(clear), 4321)
		Expect(len(ciphertext)).To(BeNumerically(">", len(clear)))
		decoded := prng.Decode(ciphertext, 4321)
		Expect(decoded).To(ContainSubstring(clear))
	})

	It("can guess a seed", func() {
		clear := "I wandered lonely as a cloud"
		ciphertext := prng.EncodeWithRandomPrefix([]byte(clear), 4321)
		seed, err := prng.GuessSeed(ciphertext, []byte(clear))
		Expect(err).NotTo(HaveOccurred())
		Expect(seed).To(Equal(uint16(4321)))
	})

	It("can guess a token seed", func() {
		email := "kieron@example.com"
		ciphertext := prng.PasswordResetToken(email)
		seed, err := prng.GuessResetTokenSeed(ciphertext, email)
		Expect(err).NotTo(HaveOccurred())
		fmt.Printf("seed = %+v\n", seed)
	})
})
