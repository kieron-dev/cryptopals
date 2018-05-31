package diffiehellman_test

import (
	"math/big"
	"sync"

	"github.com/kieron-pivotal/cryptopals/diffiehellman"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ManInTheMiddle", func() {

	var (
		client *diffiehellman.DHPeer
		server *diffiehellman.DHPeer
		mitm   *diffiehellman.ManInTheMiddle
	)

	BeforeEach(func() {
		g := big.NewInt(2)
		p := new(big.Int)
		p.SetString(bigPStr, 16)
		client = diffiehellman.NewClient(p, g)
		server = diffiehellman.NewServer()
		mitm = diffiehellman.NewManInTheMiddle(server)

	})

	It("can intercept a secure message", func() {
		cerr, serr := client.Connect(mitm)
		Expect(cerr).NotTo(HaveOccurred())
		Expect(serr).NotTo(HaveOccurred())

		var wg sync.WaitGroup
		wg.Add(2)

		var ok bool

		go func() {
			defer wg.Done()
			defer GinkgoRecover()

			msg := []byte("Hello Diffie!")
			var err error
			ok, err = client.SendTestMessage(msg)
			Expect(err).NotTo(HaveOccurred())
		}()

		go func() {
			defer wg.Done()
			mitm.ReplyTestMessage()
		}()

		wg.Wait()
		Expect(ok).To(BeTrue(), "Test message received ok")

		Expect(string(mitm.SentMsg)).To(Equal("Hello Diffie!"))
		Expect(string(mitm.ReturnedMsg)).To(Equal("Hello Diffie!"))
	})
})
