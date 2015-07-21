package main

import (
	"bytes"
	"fmt"
)

func main() {
	test := "abcdefghijklmn"
	fmt.Printf("orig: %s\nrev: %s\n", test, reverse(test))
	fmt.Printf("orig: %s\nrec rev: %s\n", test, reverse_rec(test))
}

func reverse(a string) string {
	asRune := []rune(a)
	buf := new(bytes.Buffer)

	for i := len(asRune) - 1; i >= 0; i-- {
		buf.WriteRune(asRune[i])
	}
	return buf.String()
}

func reverse_rec(a string) string {
	if len(a) < 1 {
		return ""
	}
	asRune := []rune(a)
	return reverse_rec(string(asRune[1:])) + string(asRune[0])
}
