package diffiehellman_test

import (
	"math/big"
	"sync"

	"github.com/kieron-pivotal/cryptopals/diffiehellman"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const bigPStr = "ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024" +
	"e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd" +
	"3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec" +
	"6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f" +
	"24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361" +
	"c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552" +
	"bb9ed529077096966d670c354e4abc9804f1746c08ca237327fff" +
	"fffffffffffff"

var _ = Describe("Forgot to init", func() {

	It("will exit with error if communication is not initialised prior to key exchange", func() {
		p := big.NewInt(2)
		g := big.NewInt(27)
		client := diffiehellman.NewClient(p, g)
		_, err := client.InitKeyExchange()
		Expect(err).To(HaveOccurred())

		server := diffiehellman.NewServer()
		_, err = server.CompleteKeyExchange()
		Expect(err).To(HaveOccurred())
	})
})

var _ = Describe("Normal Protocol", func() {
	var (
		sess1, sess2 []byte
		err          error
		client       *diffiehellman.DHUser
		server       *diffiehellman.DHUser
	)

	startSession := func() {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			sess1, err = client.InitKeyExchange()
			Expect(err).NotTo(HaveOccurred())
		}()

		go func() {
			defer wg.Done()
			sess2, err = server.CompleteKeyExchange()
			Expect(err).NotTo(HaveOccurred())
		}()

		wg.Wait()
	}

	BeforeEach(func() {
		g := big.NewInt(2)
		p := new(big.Int)
		p.SetString(bigPStr, 16)

		client = diffiehellman.NewClient(p, g)
		server = diffiehellman.NewServer()

		client.Connect(server)

	})

	It("can happily exchange keys", func() {
		startSession()
		Expect(sess1).To(Equal(sess2))
	})

	It("can exchange a secure message", func() {
		startSession()
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
			server.ReplyTestMessage()
		}()

		wg.Wait()
		Expect(ok).To(BeTrue(), "Test message received ok")
	})
})
