package ctr_test

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/kieron-pivotal/cryptopals/ctr"
	"github.com/kieron-pivotal/cryptopals/operations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encoder", func() {

	It("can encode and decode back to same value", func() {
		clear := "My name is Kieron"
		rand.Seed(time.Now().UnixNano())
		key := operations.RandomSlice(16)

		counter := ctr.Counter{Nonce: bytes.Repeat([]byte{0}, 8)}

		enc := ctr.Encode([]byte(clear), key, counter)
		Expect(enc).ToNot(Equal([]byte(clear)))

		decoded := ctr.Encode(enc, key, counter)
		Expect(decoded).To(Equal([]byte(clear)))
	})

	It("can edit an encrypted text", func() {
		clear := "My name is Kieron"

		rand.Seed(time.Now().UnixNano())
		key := operations.RandomSlice(16)
		counter := ctr.Counter{Nonce: operations.RandomSlice(8)}

		enc := ctr.Encode([]byte(clear), key, counter)

		edited := ctr.Edit(enc, key, 11, []byte("Ciara"), counter)

		counter.BlockCount = 0
		decoded := ctr.Encode(edited, key, counter)
		Expect(string(decoded)).To(ContainSubstring("Ciaran"))
	})

})
