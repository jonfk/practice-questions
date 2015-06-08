package main

import (
	"bytes"
	//"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/satori/go.uuid"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){}?+|=/-,.< >")

func main() {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4:\n%s\n", u1)

	passwordLength := 15
	b := make([]byte, passwordLength)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	password := new(bytes.Buffer)

	for i := range b {
		password.WriteRune(chars[int(b[i])%len(chars)])
	}
	fmt.Printf("Password:\n%s\n", password.String())
}
