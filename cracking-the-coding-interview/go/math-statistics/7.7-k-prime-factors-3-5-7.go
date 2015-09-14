package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", naive(11))
	fmt.Printf("%d\n", kPrime(11))
}

func naive(k int) int {
	var i int

	for j := 0; i < k; j++ {
		if j%3 == 0 && j%5 == 0 && j%7 == 0 {
			i++
			if i == k {
				return j
			}
		}
	}
	return -1
}

func kPrime(k int) int {
}
