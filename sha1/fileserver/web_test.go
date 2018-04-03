package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/sha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Web", func() {

	const (
		url          = "http://localhost:9000/test?file=%s&signature=%s"
		quickFox     = "The quick brown fox jumps over the lazy dog"
		quickFoxHmac = "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9"
	)

	var (
		err  error
		file *os.File
	)

	BeforeEach(func() {
		file, err = ioutil.TempFile("", "hmac")
		Expect(err).NotTo(HaveOccurred())
		_, err = file.Write([]byte(quickFox))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		file.Close()
		os.Remove(file.Name())
	})

	It("returns a file if file exists and sig is right", func() {
		get := fmt.Sprintf(url, file.Name(), quickFoxHmac)
		resp, err := http.Get(get)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(Equal([]byte(quickFox)))
	})

	DescribeTable("Empty params gets 400 error", func(urlstr string) {
		urlStr := fmt.Sprintf(url, "", "")
		resp, err := http.Get(urlStr)
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
		urlstr := fmt.Sprintf(url, file.Name(), "asdf")
		resp, err := http.Get(urlstr)
		Expect(err).NotTo(HaveOccurred())
		resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(http.StatusForbidden))
	})

	It("is quicker to get the first sig byte wrong than the second", func() {
		secondWrong := fmt.Sprintf(url, file.Name(), "de7d")
		firstWrong := fmt.Sprintf(url, file.Name(), "da7c")
		start := time.Now()
		resp, err := http.Get(secondWrong)
		end := time.Now()
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusForbidden))

		secondWrongTime := end.Sub(start)

		start = time.Now()
		resp, err = http.Get(firstWrong)
		end = time.Now()
		Expect(err).NotTo(HaveOccurred())
		resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(http.StatusForbidden))
		firstWrongTime := end.Sub(start)

		Expect(firstWrongTime).To(BeNumerically("<", secondWrongTime))
	})

	It("is possible to guess byte 1 through timing", func() {
		maxTime := time.Duration(0)
		var b byte
		for i := 0; i < 256; i++ {
			attempt := fmt.Sprintf("%x00", i)
			urlStr := fmt.Sprintf(url, file.Name(), attempt)
			t0 := time.Now()
			for j := 0; j < 6; j++ {
				resp, err := http.Get(urlStr)
				Expect(err).NotTo(HaveOccurred())
				resp.Body.Close()
			}
			t1 := time.Now()
			dur := t1.Sub(t0)
			if dur > maxTime {
				maxTime = dur
				b = byte(i)
			}
		}
		Expect(b).To(Equal(byte(0xde)))
	})

	XIt("is possible to guess the whole hash", func() {
		hash := sha1.GetSHA1HMAC(file.Name(), func(hash []byte) {
			urlStr := fmt.Sprintf(url, file.Name(), conversion.BytesToHex(hash))
			resp, err := http.Get(urlStr)
			Expect(err).NotTo(HaveOccurred())
			resp.Body.Close()
		})
		hex := conversion.BytesToHex(hash)
		Expect(hex).To(Equal(quickFoxHmac))
	})

})
