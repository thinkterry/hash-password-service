package main

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const seconds = 5
const rootPath = "/"
const shutdownPath = "/shutdown"

var srv *http.Server

func encodedHashHandler(w http.ResponseWriter, r *http.Request) {
	password, err := parsePassword(w, r)
	if err != nil {
		// HTTP status codes have already been set;
		// all that's needed to do now is:
		return
	}

	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Fprint(w, EncodedHash(password))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	validRequest := (r.URL.Path == shutdownPath && r.Method == "POST")
	if !validRequest {
		http.NotFound(w, r)
		return
	}

	StopServer()
}

func parsePassword(w http.ResponseWriter, r *http.Request) (string, error) {
	var password string

	validRequest := (r.URL.Path == rootPath && r.Method == "POST")
	if !validRequest {
		http.NotFound(w, r)
		return password, errors.New("Invalid URL or HTTP method")
	}
	err := r.ParseForm()
	if err != nil {
		badRequest(w)
		return password, errors.New("Form unparseable")
	}
	password = r.PostForm.Get("password")
	if password == "" {
		badRequest(w)
		return password, errors.New("Password paramenter missing")
	}

	return password, err
}

func badRequest(w http.ResponseWriter) {
	// per http://stackoverflow.com/a/40096757
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 bad request"))
}

func StartServer() {
	// per https://golang.org/pkg/net/http/
	http.HandleFunc(rootPath, encodedHashHandler)
	http.HandleFunc(shutdownPath, shutdownHandler)

	// per http://stackoverflow.com/a/42533360
	srv = &http.Server{Addr: ":8080"}
	go func() {
		err := srv.ListenAndServe() // block until shut down
		if err != nil {
			// probably an intentional shutdown, not an error
			log.Println(err)
		}
	}()
}

func StopServer() {
	ctx, _ := context.WithTimeout(context.Background(), seconds*time.Second)
	err := srv.Shutdown(ctx)

	if err != nil {
		// probably an intentional shutdown, not an error
		log.Println(err)
	}

	os.Exit(0) // override SIGINT blocking in main, if necessary
}

func EncodedHash(msg string) string {
	return Base64(Hash([]byte(msg)))
}

func Hash(msg []byte) []byte {
	hasher := sha512.New()
	hasher.Write(msg)
	return hasher.Sum(nil)
}

func Base64(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func main() {
	StartServer()
	defer StopServer()

	// per https://golang.org/pkg/os/signal/#example_Notify
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // block server shutdown until a SIGINT is received
}
