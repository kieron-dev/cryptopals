package examples_test

import (
	"github.com/kieron-pivotal/cryptopals/examples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CtrAttempt", func() {
	var encs [][]byte

	BeforeEach(func() {
		encs = examples.EncryptList("../assets/03_19.txt")
	})

	It("can encrypt the list", func() {
		Expect(len(encs)).To(BeNumerically(">", 0))
	})
})
