package examples

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kieron-pivotal/cryptopals/operations"
)

const blocksize = 16

var bitFlipKey = operations.RandomSlice(blocksize)
var bitFlipIV = operations.RandomSlice(blocksize)

const prefix = "comment1=cooking%20MCs;userdata="
const suffix = ";comment2=%20like%20a%20pound%20of%20bacon"

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
