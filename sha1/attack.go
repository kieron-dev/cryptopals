package sha1

import (
	"bytes"
	"fmt"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
)

func GetSHA1HMAC(property string, makeCall func(hash []byte)) []byte {
	hashLen := 20
	hash := bytes.Repeat([]byte{0}, hashLen)

	for i := 0; i < hashLen; i++ {
		crackByte(hash, i, property, makeCall)
		fmt.Println(conversion.BytesToHex(hash))
	}

	return hash
}

func crackByte(hash []byte, pos int, prop string, makeCall func(hash []byte)) {
	maxTime := time.Duration(0)
	var b byte
	for i := 0; i < 256; i++ {
		hash[pos] = byte(i)
		t0 := time.Now()
		for j := 0; j < 8; j++ {
			makeCall(hash)
		}
		t1 := time.Now()
		dur := t1.Sub(t0)
		if dur > maxTime {
			maxTime = dur
			b = byte(i)
		}
	}
	hash[pos] = b
}
