package operations

import (
	"math/rand"
	"time"
)

func RandomSlice(l int) []byte {
	rand.Seed(time.Now().UnixNano())
	slice := []byte{}
	for i := 0; i < l; i++ {
		slice = append(slice, byte(rand.Intn(256)))
	}
	return slice
}
