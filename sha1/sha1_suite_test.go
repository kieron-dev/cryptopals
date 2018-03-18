package sha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSha1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SHA1 Suite")
}
