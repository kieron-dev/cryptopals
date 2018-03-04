package examples

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/kieron-pivotal/cryptopals/operations"
)

var (
	kvKey []byte
)

func init() {
	rand.Seed(time.Now().UnixNano())
	kvKey = operations.RandomSlice(16)
}

func ParseKVString(s string) map[string]string {
	ret := map[string]string{}

	for _, pair := range strings.Split(s, "&") {
		terms := strings.Split(pair, "=")
		ret[terms[0]] = terms[1]
	}
	return ret
}

func EncodeKVs(kvs map[string]string) string {
	pairs := []string{}
	var keys []string
	for k := range kvs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	if email, ok := kvs["email"]; ok {
		pairs = append(pairs, fmt.Sprintf("email=%s", email))
	}

	for _, k := range keys {
		if k == "email" || k == "role" {
			continue
		}
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, kvs[k]))
	}

	if role, ok := kvs["role"]; ok {
		pairs = append(pairs, fmt.Sprintf("role=%s", role))
	}

	return strings.Join(pairs, "&")
}

func ProfileFor(email string) map[string]string {
	email = strings.Replace(email, "&", "", -1)
	email = strings.Replace(email, "=", "", -1)
	return map[string]string{
		"email": email,
		"uid":   "10",
		"role":  "user",
	}
}

func GetCookie(email string) []byte {
	profile := ProfileFor(email)
	str := EncodeKVs(profile)
	enc, err := operations.AES128ECBEncode([]byte(str), kvKey)
	if err != nil {
		panic(err)
	}
	return enc
}

func DecryptCookie(cookie []byte) map[string]string {
	clear, err := operations.AES128ECBDecode(cookie, kvKey)
	clear = operations.RemovePKCS7(clear, 16)
	if err != nil {
		panic(err)
	}
	return ParseKVString(string(clear))
}

func GetAdminCookie() []byte {
	elevens := bytes.Repeat([]byte{11}, 11)
	cookieWithAdminBlock := GetCookie("foo@bar.coadmin" + string(elevens))
	cookieWithRoleInOwnBlock := GetCookie("foo01@bar.com")
	l := len(cookieWithRoleInOwnBlock)
	for i := 0; i < 16; i++ {
		cookieWithRoleInOwnBlock[l-16+i] = cookieWithAdminBlock[16+i]
	}
	return cookieWithRoleInOwnBlock
}
