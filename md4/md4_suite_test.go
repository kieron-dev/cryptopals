package md4_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMd4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Md4 Suite")
}
