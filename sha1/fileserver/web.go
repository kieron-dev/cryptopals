package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kieron-pivotal/cryptopals/conversion"
	"github.com/kieron-pivotal/cryptopals/sha1"
)

var key = []byte("key")

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
	if hexHMAC != sig {
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
