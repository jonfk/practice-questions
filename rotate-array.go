package main

import (
	"fmt"
)

/*
Problem:
Rotate an array of n elements to the right by k steps.
For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7]
is rotated to [5,6,7,1,2,3,4]. How many different ways do you know to solve this problem?
*/

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v", slice)
	fmt.Printf("%v", rotate(slice, 3))
}

func rotate(slice []int, k int) []int {
	if len(slice) < k {
		return slice
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if i+k <= len(slice)-1 {
			slice[i], slice[i+k] = slice[i+k], slice[i]
		}
	}
	return slice
}
