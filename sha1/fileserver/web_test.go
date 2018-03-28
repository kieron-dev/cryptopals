package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Web", func() {

	var (
		session *gexec.Session
		err     error
		url     string
	)

	BeforeEach(func() {
		command := exec.Command(pathToExe)
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())
		url = "http://localhost:9000/test?file=%s&signature=%s"
	})

	It("can start the web server", func() {
		Eventually(session.Out).Should(gbytes.Say("Listening"))
	})

	It("returns a file if file exists and sig is right", func() {
		file, err := ioutil.TempFile("", "hmac")
		Expect(err).NotTo(HaveOccurred())

		_, err = file.Write([]byte("The quick brown fox jumps over the lazy dog"))
		Expect(err).NotTo(HaveOccurred())
		defer os.Remove(file.Name())
		defer file.Close()

		get := fmt.Sprintf(url, file.Name(), "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9")
		resp, err := http.Get(get)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())

		Expect(body).To(Equal([]byte("The quick brown fox jumps over the lazy dog")))
	})

	DescribeTable("Empty params gets 400 error", func(urlstr string) {
		resp, err := http.Get(url)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
	},
		Entry("empty sig", fmt.Sprintf(url, "a", "")),
		Entry("empty file", fmt.Sprintf(url, "", "b")),
		Entry("empty both", fmt.Sprintf(url, "", "")),
	)

	It("if file not found get 404", func() {
		urlstr := fmt.Sprintf(url, "/does/not/exist", "asdf")
		resp, err := http.Get(urlstr)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
	})

	It("if sig wrong get 403", func() {
		file, err := ioutil.TempFile("", "hmac")
		Expect(err).NotTo(HaveOccurred())

		_, err = file.Write([]byte("The quick brown fox jumps over the lazy dog"))
		Expect(err).NotTo(HaveOccurred())
		defer os.Remove(file.Name())
		defer file.Close()

		urlstr := fmt.Sprintf(url, file.Name(), "asdf")
		resp, err := http.Get(urlstr)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusForbidden))
	})

})
