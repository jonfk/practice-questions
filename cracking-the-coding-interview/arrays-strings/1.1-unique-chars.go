package main

import (
	"fmt"
)

func duplicates(a string) bool {
	var set map[rune]int = make(map[rune]int)
	b := []rune(a)
	for i := range b {
		if set[b[i]] > 0 {
			return true
		} else {
			set[b[i]] = 1
		}
	}
	return false
}

func main() {
	dup := "aaaoeub"
	nodup := "aoeu',.p"
	fmt.Printf("true: %v\nfalse: %v\n", duplicates(dup), duplicates(nodup))
}
