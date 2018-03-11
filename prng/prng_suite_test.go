package prng_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrng(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prng Suite")
}
