package main

import (
	"fmt"
)

/*
Given 2 integer arrays, determine of the 2nd array is a rotated version of the 1st array.
Ex. Original Array A={1,2,3,5,6,7,8} Rotated Array B={5,6,7,8,1,2,3}
*/

func main() {
	//a, b := []int{1, 2, 3, 5, 6, 7, 8}, []int{5, 6, 7, 8, 1, 2, 3}
	a, b := []int{1, 2, 3, 5, 6, 7, 8}, []int{1, 2, 3, 5, 6, 7, 8}
	fmt.Printf("%v\n", isRotated(a, b))
}

func isRotated(a, b []int) bool {
	i, j := 0, 0
	for ; j < len(b); j++ {
		if a[i] == b[j] {
			break
		}
	}
	if j == len(a) {
		return false
	}
	for j < len(b) && i < len(a) && a[i] == b[j] {
		i++
		j++
	}
	if i == len(a) && j == len(b) {
		return true
	} else if i == len(a) {
		i = 0
		for j < len(b) && i < len(a) && a[i] == b[j] {
			i++
			j++
		}
		if j == len(b) {
			return true
		} else {
			return false
		}
	} else if j == len(b) {
		j = 0
		for j < len(b) && i < len(a) && a[i] == b[j] {
			i++
			j++
		}
		if i == len(a) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}
