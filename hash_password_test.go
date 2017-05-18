package main

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestHash(t *testing.T) {
	hashed := Hash([]byte{'a'})
	expected, _ := hex.DecodeString("" +
		"1f40fc92da241694750979ee6cf582f2" +
		"d5d7d28e18335de05abc54d0560e0f53" +
		"02860c652bf08d560252aa5e74210546" +
		"f369fbbbce8c12cfc7957b2652fe9a75")
	if !bytes.Equal(hashed, expected) {
		t.Errorf("Expected:\n%v\nbut got:\n%v\n", expected, hashed)
	}
}

func TestBase64(t *testing.T) {
	encoded := Base64([]byte{'a'})
	const expected = "YQ=="
	if encoded != expected {
		t.Errorf("Expected:\n%v\nbut got:\n%v\n", expected, encoded)
	}
}

func TestEncodedHash(t *testing.T) {
	encoded := EncodedHash("angryMonkey")
	const expected = "" +
		"ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aT" +
		"aRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0v" +
		"KbZPZklJz0Fd7su2A+gf7Q=="
	if encoded != expected {
		t.Errorf("Expected:\n%v\nbut got:\n%v\n", expected, encoded)
	}
}

// end-to-end test
func TestService(t *testing.T) {
	StartServer()
	defer StopServer()

	data := url.Values{}
	data.Set("password", "angryMonkey")
	r, err := http.PostForm("http://localhost:8080/", data)
	if err != nil {
		t.Error(err)
	}
	encodedArr, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		t.Error(err)
	}
	encoded := string(encodedArr)
	const expected = "" +
		"ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aT" +
		"aRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0v" +
		"KbZPZklJz0Fd7su2A+gf7Q=="
	if encoded != expected {
		t.Errorf("Expected:\n%v\nbut got:\n%v\n", expected, encoded)
	}
}
