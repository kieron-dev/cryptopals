package ctr

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/kieron-pivotal/cryptopals/operations"
)

var (
	attackKey   []byte
	attackNonce []byte
)

func init() {
	attackKey = operations.RandomSlice(16)
	attackNonce = operations.RandomSlice(8)
}

func ExampleToken(userdata string) []byte {
	userdata = strings.Replace(userdata, ";", "", -1)
	userdata = strings.Replace(userdata, "=", "", -1)
	clear := fmt.Sprintf("comment1=cooking%%20MCs;userdata=%s;comment2=%%20like%%20a%%20pound%%20of%%20bacon", userdata)

	c := Counter{Nonce: attackNonce}
	return Encode([]byte(clear), attackKey, c)
}

func CheckForAdmin(enc []byte) bool {
	c := Counter{Nonce: attackNonce}
	clear := Encode(enc, attackKey, c)
	return bytes.Contains(clear, []byte(";admin=true;"))
}

func GetVarInputPos() (int, error) {
	enc1 := ExampleToken("000000000000")
	enc2 := ExampleToken("111111111111")
	for i := 0; i < len(enc1); i++ {
		if enc1[i] != enc2[i] {
			return i, nil
		}
	}
	return 0, errors.New("couldn't determine position")
}
