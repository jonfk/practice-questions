package main

import (
	"fmt"
)

type Result struct {
	A, B int
}

func (a Result) Equals(b Result) bool {
	if (a.A == b.A && a.B == b.B) || (a.B == b.A && a.A == b.B) {
		return true
	}
	return false
}

type Results []Result

func (a Results) Exists(b Result) bool {
	for i := range a {
		if a[i].Equals(b) {
			return true
		}
	}
	return false
}

func main() {
	test := []int{9, 1, 8, 8, 7, 6, 5, 7, 6, 4, 32, 3, 45, 6, 11, 8, 9, 7, 2, 78, 9, 0, 10}
	fmt.Printf("%v\n", pairs(test, 10))
}

func pairs(a []int, sum int) Results {
	var result Results
	for i := range a {
		complement := sum - a[i]
		if ExistsInt(a, complement, i) {
			candidate := Result{a[i], complement}
			if !result.Exists(candidate) {
				result = append(result, candidate)
			}
		}
	}
	return result
}

func ExistsInt(a []int, x, except int) bool {
	for i := range a {
		if a[i] == x && i != except {
			return true
		}
	}
	return false
}
