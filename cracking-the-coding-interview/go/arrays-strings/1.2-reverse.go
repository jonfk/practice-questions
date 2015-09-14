package main

import (
	"fmt"
)

func reverse(a string) string {
	var result []rune
	b := []rune(a)
	for i := range b {
		result = append([]rune{b[i]}, result...)
	}
	return string(result)
}

func main() {
	test := "abcdefghijklhm"
	fmt.Printf("%s\n", reverse(test))
}
