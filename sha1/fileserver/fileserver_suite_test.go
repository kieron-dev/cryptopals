package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestFileserver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fileserver Suite")
}

var pathToExe string

var _ = BeforeSuite(func() {
	var err error
	pathToExe, err = gexec.Build("github.com/kieron-pivotal/cryptopals/sha1/fileserver")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
