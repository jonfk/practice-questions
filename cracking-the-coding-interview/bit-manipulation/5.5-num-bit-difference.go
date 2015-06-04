package main

import (
	"fmt"
)

func isSet(num, i int) bool {
	return (num & (1 << uint(i))) != 0
}

func numSetBits(a int) int {
	var setBits int
	for i := 0; i < 32; i++ {
		if isSet(a, i) {
			setBits++
		}
	}
	return setBits
}

func difference(a, b int) int {
	one := a | b
	zero := (^a) | (^b)
	difference := one & zero
	return numSetBits(difference)
}

func main() {
	a := 31
	b := 14
	fmt.Printf("a: %b\nb: %b\n", a, b)
	fmt.Printf("differenc: %d\n", difference(a, b))
}
