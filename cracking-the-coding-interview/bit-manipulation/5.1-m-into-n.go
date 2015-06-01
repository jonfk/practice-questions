package main

import (
	"fmt"
)

func mIntoN(n, m, i, j int) int {
	result := n
	var mi, x uint

	for x = uint(i); x <= uint(j); x++ {
		bitInM := ((1 << mi) & m) != 0
		var mBit int
		if bitInM {
			mBit = 1
		} else {
			mBit = 0
		}
		result = (result & ^(1 << x)) | (mBit << x)
		mi++
	}
	return result
}

func main() {
	n := (1 << 10)
	m := 19
	fmt.Printf("n: %b\n", n)
	fmt.Printf("m: %b\n", m)
	fmt.Printf("result : %b\n", mIntoN(n, m, 2, 6))
}
