package md4_test

import (
	"github.com/kieron-pivotal/cryptopals/md4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MD4", func() {
	It("can compute an md4 sum", func() {
		sum := md4.Sum([]byte("foo bar"))
		Expect(sum).To(Equal("2923f5cdcd3c485e73413d92cf26839b"))
	})
})
