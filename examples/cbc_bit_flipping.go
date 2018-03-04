package examples

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kieron-pivotal/cryptopals/operations"
)

const (
	blocksize = 16
	prefix    = "comment1=cooking%20MCs;userdata="
	suffix    = ";comment2=%20like%20a%20pound%20of%20bacon"
)

var (
	bitFlipKey []byte
	bitFlipIV  []byte
)

func init() {
	rand.Seed(time.Now().UnixNano())
	bitFlipKey = operations.RandomSlice(blocksize)
	bitFlipIV = operations.RandomSlice(blocksize)
}

func EncodeUserdata(userdata string) []byte {
	userdata = strings.Replace(userdata, ";", "", -1)
	userdata = strings.Replace(userdata, "=", "", -1)

	toEncode := fmt.Sprintf("%s%s%s", prefix, userdata, suffix)
	clear := operations.PKCS7([]byte(toEncode), blocksize)

	enc, err := operations.AES128CBCEncode(clear, bitFlipKey, bitFlipIV)
	if err != nil {
		panic(err)
	}

	return enc
}

func IsAdmin(enc []byte) bool {
	clear, err := operations.AES128CBCDecode(enc, bitFlipKey, bitFlipIV)
	if err != nil {
		panic(err)
	}

	return bytes.Contains(clear, []byte(";admin=true;"))
}
