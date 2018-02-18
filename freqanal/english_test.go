package freqanal_test

import (
	"github.com/kieron-pivotal/cryptopals/freqanal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("English", func() {

	DescribeTable("GetFreqs",
		func(in string, freqs map[rune]int) {

			Expect(freqanal.GetFreqs(in)).
				To(Equal(freqs))
		},
		Entry("foo", "foo", map[rune]int{'f': 1, 'o': 2}),
		Entry("xxx", "xxx", map[rune]int{'x': 3}),
	)

	DescribeTable("FreqScoreStr",
		func(in, out string) {
			Expect(freqanal.FreqSortStr(in)).To(Equal(out))
		},

		Entry("foo", "foo", "of"),
		Entry("FoO", "FoO", "of"),
		Entry(" Fo O", "Fo O", "of"),
	)

	DescribeTable("FreqScore",
		func(in string, score int) {
			Expect(freqanal.FreqScoreEnglish(in)).To(Equal(score))
		},

		Entry("empty", "", 0),
		Entry("first 3 correct position", "eeetta", 60),
	)

})
