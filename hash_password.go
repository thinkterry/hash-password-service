package main

import (
	"crypto/sha512"
	"fmt"
)

func Hash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
