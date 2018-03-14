package prng

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/operations"
)

func Encode(clear []byte, seed uint32) []byte {
	mt := New(uint32(seed))
	ret := []byte{}

	for i := 0; i < len(clear); i++ {
		num := mt.Next()
		ret = append(ret, clear[i]^uint8(num))
	}

	return ret
}

func Decode(ciphertext []byte, seed uint32) []byte {
	return Encode(ciphertext, seed)
}

func EncodeWithRandomPrefix(clear []byte, seed uint16) []byte {
	rand.Seed(time.Now().UnixNano())
	prefixLen := rand.Intn(30) + 1
	prefix := operations.RandomSlice(prefixLen)
	return Encode(append(prefix, clear...), uint32(seed))
}

func GuessSeed(ciphertext, knownClear []byte) (uint16, error) {
	for i := 0; i < 1<<16; i++ {
		decoded := Decode(ciphertext, uint32(i))
		if bytes.Contains(decoded, knownClear) {
			return uint16(i), nil
		}
	}
	return 0, fmt.Errorf("Couldn't guess seed")
}

func PasswordResetToken(email string) []byte {
	clear := fmt.Sprintf("id=2431&email=%s&timestamp=123532434", email)
	return Encode([]byte(clear), uint32(time.Now().Unix()))
}

func GuessResetTokenSeed(ciphertext []byte, email string) (uint32, error) {
	t := uint32(time.Now().Unix())
	for i := 0; i < 3600; i++ {
		s := t - uint32(i)
		decoded := Decode(ciphertext, s)
		if bytes.Contains(decoded, []byte(email)) {
			return s, nil
		}
	}
	return 0, errors.New("couldn't find token seed")
}
