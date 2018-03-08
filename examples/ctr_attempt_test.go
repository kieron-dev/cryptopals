package examples_test

import (
	"fmt"
	"strings"

	"github.com/kieron-pivotal/cryptopals/examples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CtrAttempt", func() {
	var encs [][]byte

	BeforeEach(func() {
		encs = examples.EncryptList()
	})

	It("can encrypt the list", func() {
		fmt.Println(encs)
		Expect(len(encs)).To(BeNumerically(">", 0))
	})

	FIt("can attempt a decryption", func() {
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
		fmt.Println()
		fmt.Println(strings.Repeat("0 1 2 3 4 5 6 7 8 9 ", 3))
		printDecryption(encs, attempt)
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
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == ' ' ||
		char == '\n' ||
		char == ',' ||
		char == '.' ||
		char == '\'' ||
		char == '!' ||
		char == '-'
}
