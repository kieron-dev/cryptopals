package diffiehellman_test

import (
	"math/big"
	"sync"

	"github.com/kieron-pivotal/cryptopals/diffiehellman"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Protocol", func() {
	var (
		sess1, sess2 []byte
		ch           chan string
	)

	BeforeEach(func() {
		g := big.NewInt(2)
		p := new(big.Int)
		p.SetString("ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024"+
			"e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd"+
			"3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec"+
			"6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f"+
			"24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361"+
			"c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552"+
			"bb9ed529077096966d670c354e4abc9804f1746c08ca237327fff"+
			"fffffffffffff", 16)

		ch = make(chan string)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			sess1 = diffiehellman.StartSession(p, g, ch)
			wg.Done()
		}()

		go func() {
			sess2 = diffiehellman.PairSession(ch)
			wg.Done()
		}()

		wg.Wait()
	})

	It("can happily exchange keys", func() {
		Expect(sess1).To(Equal(sess2))
	})

	It("can exchange a secure message", func() {
		var wg sync.WaitGroup
		wg.Add(2)

		var ok bool

		go func() {
			msg := []byte("Hello Diffie!")
			var err error
			ok, err = diffiehellman.SendTestMessage(ch, sess1, msg)
			Expect(err).NotTo(HaveOccurred())
			wg.Done()
		}()

		go func() {
			diffiehellman.ReplyTestMessage(ch, sess2)
			wg.Done()
		}()

		wg.Wait()
		Expect(ok).To(BeTrue(), "Test message received ok")
	})
})
