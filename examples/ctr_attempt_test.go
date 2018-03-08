package examples_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/examples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CtrAttempt", func() {
	var encs [][]byte

	BeforeEach(func() {
		encs = examples.EncryptList()
	})

	It("can encrypt the list", func() {
		fmt.Println(encs)
		Expect(len(encs)).To(BeNumerically(">", 0))
	})
})
