package main

import (
	"fmt"
)

func main() {
	test := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

	fmt.Printf("%v %d\n", test, stringLoc(test, "f", 0, len(test)-1))
}

func stringLoc(a []string, b string, start, end int) int {
	if end <= start || end >= len(a) || start < 0 {
		return -1
	}

	mid := (end + start) / 2
	firstBlank := mid
	for a[mid] == "" {
		mid++
	}

	if a[mid] == b {
		return mid
	} else if a[mid] > b {
		return stringLoc(a, b, start, firstBlank)
	} else {
		return stringLoc(a, b, mid, end)
	}

}
