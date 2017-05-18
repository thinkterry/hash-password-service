package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
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
	validRequest := (r.URL.Path == "/" && r.Method == "POST")
	if !validRequest {
		http.NotFound(w, r)
		return
	}
	err := r.ParseForm()
	if err != nil {
		badRequest(w)
		return
	}
	password := r.PostForm.Get("password")
	if password == "" {
		badRequest(w)
		return
	}
	fmt.Fprint(w, EncodedHash(password))
}

func badRequest(w http.ResponseWriter) {
	// per http://stackoverflow.com/a/40096757
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 bad request"))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	validRequest := (r.URL.Path == "/shutdown" && r.Method == "POST")
	if !validRequest {
		http.NotFound(w, r)
		return
	}
	log.Fatal("Not-so-graceful shutdown")
}

func main() {
	// per https://golang.org/pkg/net/http/
	http.HandleFunc("/", handler)
	http.HandleFunc("/shutdown", shutdownHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
