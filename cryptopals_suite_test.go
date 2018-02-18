package cryptopals_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCryptopals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cryptopals Suite")
}
