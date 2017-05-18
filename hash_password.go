package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
)

func Hash(msg []byte) []byte {
	hasher := sha512.New()
	hasher.Write(msg)
	return hasher.Sum(nil)
}

func Base64(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func EncodedHash(msg string) string {
	return Base64(Hash([]byte(msg)))
}

// per https://golang.org/doc/articles/wiki/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
