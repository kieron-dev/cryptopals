package freqanal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFreqanal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Freqanal Suite")
}
