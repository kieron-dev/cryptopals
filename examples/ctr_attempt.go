package examples

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/ctr"
	"github.com/kieron-pivotal/cryptopals/operations"
)

var (
	ctrKey  []byte
	counter ctr.Counter
)

func init() {
	rand.Seed(time.Now().UnixNano())
	ctrKey = operations.RandomSlice(16)
	counter = ctr.Counter{Nonce: operations.RandomSlice(8)}
}

func EncryptList() [][]byte {
	file, err := os.Open("assets/03_19.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encs := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter.BlockCount = 0
		clear, err := conversion.Base64ToBytes(scanner.Text())
		if err != nil {
			panic(err)
		}
		encs = append(encs, ctr.Encode(clear, ctrKey, counter))
	}
	return encs
}
