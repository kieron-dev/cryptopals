package examples_test

import (
	"github.com/kieron-pivotal/cryptopals/examples"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CbcBitFlipping", func() {
	It("isn't possible to get admin by passing it in", func() {
		enc := examples.EncodeUserdata(";admin=true;")
		Expect(examples.IsAdmin(enc)).ToNot(BeTrue())
	})
})
