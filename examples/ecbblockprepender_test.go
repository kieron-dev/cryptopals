package examples_test

import (
	"fmt"

	"github.com/kieron-pivotal/cryptopals/examples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ECBBlockPrepender", func() {

	It("finds the first letter", func() {
		ret := examples.ECBBlockPrependerOracle(examples.ECBBlockPrependerEncode)
		fmt.Println(string(ret))
		Expect(len(ret)).To(BeNumerically(">", 0))
	})

})
