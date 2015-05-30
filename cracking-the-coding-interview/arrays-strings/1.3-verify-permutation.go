package main

import (
	"fmt"
)

func main() {
	a, b := "aaoeu", "uaeoa"
	c, d := "tehuoes", "tehur,"
	fmt.Printf("true: %v\nfalse: %v\n", checkPermutation(a, b), checkPermutation(c, d))
}

func checkPermutation(a, b string) bool {
	var hasha, hashb map[rune]int = make(map[rune]int), make(map[rune]int)

	a1 := []rune(a)
	b1 := []rune(b)
	for i := range a1 {
		hasha[a1[i]] = hasha[a1[i]] + 1
	}
	for i := range b1 {
		hashb[b1[i]] = hashb[b1[i]] + 1
	}
	fmt.Printf("hasha : %v\nhashb : %v\n", hasha, hashb)
	for key, val := range hasha {
		if val != hashb[key] {
			return false
		}
	}
	for key, val := range hashb {
		if val != hasha[key] {
			return false
		}
	}
	return true
}
