package operations

import (
	"math/rand"
)

func RandomSlice(l int) []byte {
	slice := []byte{}
	for i := 0; i < l; i++ {
		slice = append(slice, byte(rand.Intn(256)))
	}
	return slice
}
