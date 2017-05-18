package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
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

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	password := string(data)
	fmt.Fprint(w, EncodedHash(password))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Graceful shutdown.")
}

func main() {
	// per https://golang.org/pkg/net/http/
	http.HandleFunc("/", handler)
	http.HandleFunc("/shutdown", shutdownHandler)
	http.ListenAndServe(":8080", nil)
}
