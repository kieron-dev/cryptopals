package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/sha1"
)

var (
	key           = []byte("key")
	sleepInterval = 5 * time.Millisecond
)

func main() {
	http.HandleFunc("/test", testHandler)
	fmt.Fprintln(os.Stdout, "Listening on http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	file := query.Get("file")
	sig := query.Get("signature")
	if file == "" || sig == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "<h1>Bad Request</h1>")
		return
	}

	hmac, err := sha1.FileHMAC(key, file)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "<h1>Not Found</h1>")
		return
	}

	hexHMAC := conversion.BytesToHex(hmac)
	if !bytesAreEqual(hexHMAC, sig) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "<h1>Forbidden</h1>")
		return
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "<h1>Not Found</h1>")
		return
	}

	w.Write(b)
}

func bytesAreEqual(hmac, sig string) bool {
	hmacBytes, err := conversion.HexToBytes(hmac)
	if err != nil {
		return false
	}
	sigBytes, err := conversion.HexToBytes(sig)
	if err != nil {
		return false
	}

	for i := 0; i < len(hmacBytes) && i < len(sigBytes); i++ {
		if hmacBytes[i] != sigBytes[i] {
			return false
		}
		time.Sleep(sleepInterval)
	}
	return len(hmacBytes) == len(sigBytes)
}
