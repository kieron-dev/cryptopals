package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestFileserver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fileserver Suite")
}

var _ = BeforeSuite(func() {
	var err error
	pathToExe, err := gexec.Build("github.com/kieron-pivotal/cryptopals/sha1/fileserver")
	Expect(err).NotTo(HaveOccurred())
	command := exec.Command(pathToExe)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ShouldNot(HaveOccurred())
	Eventually(session).Should(gbytes.Say("Listening"))

})

var _ = AfterSuite(func() {
	gexec.Terminate()
	gexec.CleanupBuildArtifacts()
})
