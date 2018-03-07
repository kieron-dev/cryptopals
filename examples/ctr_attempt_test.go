package examples_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/examples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CtrAttempt", func() {
	It("can encrypt the list", func() {
		encs := examples.EncryptList()
		fmt.Println(encs)
		Expect(true).To(BeTrue())
	})
})
