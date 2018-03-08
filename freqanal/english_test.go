package freqanal_test

import (
	"github.com/kieron-pivotal/cryptopals/freqanal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("English", func() {

	It("scores on letter frequency", func() {
		normalScore := freqanal.FreqScoreEnglish([]byte("This is a normal sentence"))
		forcedScore := freqanal.FreqScoreEnglish([]byte("The quick brown dogs jumped over the lazy fox"))
		randomScore := freqanal.FreqScoreEnglish([]byte("&6%4dsfhk223sdoi s dfhjsfdl12*&ydas"))

		Expect(normalScore).To(BeNumerically("<", forcedScore))
		Expect(forcedScore).To(BeNumerically("<", randomScore))
	})

})
