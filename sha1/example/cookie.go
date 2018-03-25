package example

import (
	"math/rand"
	"strings"
	"time"

	"github.com/kieron-pivotal/cryptopals/operations"
	"github.com/kieron-pivotal/cryptopals/sha1"
)

var key []byte

func init() {
	rand.Seed(time.Now().UnixNano())
	keyLen := 16 + rand.Intn(32)
	key = operations.RandomSlice(keyLen)
}

func GenerateCookie() (clear, hash string) {
	clear = "comment1=cooking%20MCs;userdata=foo;comment2=%20like%20a%20pound%20of%20bacon"
	hash = sha1.GenerateSHA1MAC(key, []byte(clear))
	return
}

func VerifyCookie(clear, hash string) bool {
	return sha1.VerifySHA1MAC(hash, key, []byte(clear))
}

func IsAdmin(clear, hash string) bool {
	if !VerifyCookie(clear, hash) {
		return false
	}
	return strings.Contains(clear, ";admin=true")
}
