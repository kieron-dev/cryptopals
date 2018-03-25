package md4_test

import (
	"io"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/md4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MD4", func() {
	It("can compute an md4 sum", func() {
		d := md4.New()
		in := "foo bar"
		io.WriteString(d, in)
		sum := d.Sum(nil)
		sumHex := conversion.BytesToHex(sum)
		Expect(sumHex).To(Equal("2923f5cdcd3c485e73413d92cf26839b"))
	})
})
