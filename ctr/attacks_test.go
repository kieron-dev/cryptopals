package ctr_test

import (
	"github.com/kieron-pivotal/cryptopals/ctr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attacks", func() {
	It("won't detect admin=true if passed in directly", func() {
		enc := ctr.ExampleToken(";admin=true;")
		Expect(ctr.CheckForAdmin(enc)).To(BeFalse())
	})

	It("can detect where my input goes", func() {
		fromPos, err := ctr.GetVarInputPos()
		Expect(err).NotTo(HaveOccurred())
		Expect(fromPos).To(Equal(32))
	})
})
