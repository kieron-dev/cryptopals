package cryptopals_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/ctr"
	"github.com/kieron-pivotal/cryptopals/examples"
	"github.com/kieron-pivotal/cryptopals/operations"
	"github.com/kieron-pivotal/cryptopals/prng"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CryptopalsSet03", func() {

	It("question 17", func() {
		enc, iv := examples.EncodeRandomText()
		clear := examples.PaddingOracle(enc, iv, examples.IsCorrectlyPadded)

		fmt.Println("---")
		fmt.Println(string(operations.RemovePKCS7(clear, 16)))
		fmt.Println("---")

		Expect(clear).To(HaveLen(len(enc)))
		Expect(string(clear)).To(MatchRegexp(`^0000`))
	})

	It("question 18", func() {
		enc64 := "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
		enc, err := conversion.Base64ToBytes(enc64)
		if err != nil {
			panic(err)
		}

		key := []byte("YELLOW SUBMARINE")
		counter := ctr.Counter{
			Nonce: bytes.Repeat([]byte{0}, 8),
		}

		clear := ctr.Encode(enc, key, counter)
		fmt.Println(string(clear))

		Expect(string(clear)).To(ContainSubstring("Ice, Ice, baby"))
	})

	It("question 19", func() {
		encs := examples.EncryptList("./assets/03_19.txt")

		attempt := initAttempt(encs)

		for i, _ := range attempt[0] {
			attempt[0][i] = 'v'
		}
		attempt[0][0] = 'i'
		attempt[0][1] = ' '
		attempt[0][2] = 'h'
		attempt[0][3] = 'a'
		attempt[0][4] = 'v'
		attempt[0][5] = 'e'
		attempt[0][6] = ' '
		attempt[0][7] = 'm'
		attempt[0][8] = 'e'
		attempt[0][9] = 't'
		attempt[0][10] = ' '
		attempt[0][11] = 't'
		attempt[0][12] = 'h'
		attempt[0][13] = 'e'
		attempt[0][14] = 'm'
		attempt[0][15] = ' '
		attempt[0][16] = 'a'
		attempt[0][17] = 't'
		attempt[0][18] = ' '
		attempt[0][19] = 'c'
		attempt[0][20] = 'l'
		attempt[0][21] = 'o'
		attempt[0][22] = 's'
		attempt[0][23] = 'e'
		attempt[0][24] = ' '
		attempt[0][25] = 'o'
		attempt[2][26] = ' '
		attempt[6][27] = ' '
		attempt[0][28] = 'd'
		attempt[0][29] = 'a'
		attempt[0][30] = 'y'
		attempt[6][31] = 'd'
		attempt[4][32] = 'h'
		attempt[4][33] = 'e'
		attempt[4][34] = 'a'
		attempt[4][35] = 'd'
		attempt[37][36] = 'n'
		attempt[37][37] = ','
		fmt.Println()
		fmt.Println(strings.Repeat("0 1 2 3 4 5 6 7 8 9 ", 3))
		printDecryption(encs, attempt)
	})

	It("question 20", func() {
		encs := examples.EncryptList("./assets/03_20.txt")

		minLength := 10000
		for _, r := range encs {
			if len(r) < minLength {
				minLength = len(r)
			}
		}

		buf := []byte{}
		for _, r := range encs {
			buf = append(buf, r[:minLength]...)
		}

		clear, _ := operations.RepeatingXorDecrypt(buf)
		Expect(clear).To(ContainSubstring("the knowledge"))

		for len(clear) > 0 {
			fmt.Println(clear[:minLength])
			clear = clear[minLength:]
		}
	})

	XIt("question 22", func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(960)+40) * time.Second)
		rng := prng.New(uint32(time.Now().Unix()))
		r := rng.Next()
		time.Sleep(time.Duration(rand.Intn(960)+40) * time.Second)

		now := uint32(time.Now().Unix())

		foundIt := false
		for i := uint32(0); i < 1100; i++ {
			seed := now - i
			rnd := prng.New(seed)
			if rnd.Next() == r {
				foundIt = true
				fmt.Printf("Seed was %d\n", seed)
				break
			}
		}

		Expect(foundIt).To(BeTrue(), "found the seed")
	})

	It("question 23 - cloning a MT19937", func() {
		rand.Seed(time.Now().UnixNano())
		mer1 := prng.New(rand.Uint32())
		var seed []uint32
		for i := 0; i < 624; i++ {
			seed = append(seed, prng.Detemper(mer1.Next()))
		}
		mer2 := prng.Mersenne{}
		mer2.Seed(seed)

		for i := 0; i < 1000; i++ {
			Expect(mer2.Next()).To(Equal(mer1.Next()))
		}
	})
})

func initAttempt(enc [][]byte) [][]byte {
	attempt := [][]byte{}
	for _, r := range enc {
		attempt = append(attempt, make([]byte, len(r)))
	}
	return attempt
}

func printDecryption(enc, guessedClear [][]byte) {
	streamLen := 0
	for _, r := range enc {
		if len(r) > streamLen {
			streamLen = len(r)
		}
	}
	stream := make([]byte, streamLen)

	for j, r := range guessedClear {
		for i, c := range r {
			if c != byte(0) {
				stream[i] = enc[j][i] ^ c
			}
		}
	}

	for _, r := range enc {
		line := ""
		for i, c := range r {
			char := stream[i] ^ c
			if isPrintable(char) {
				line += string(char)
			} else {
				line += "*"
			}
			line += " "
		}
		fmt.Println(line)
	}
}

func isPrintable(char byte) bool {
	return (char >= 32 && char <= 126)
}
