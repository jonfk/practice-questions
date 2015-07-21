package main

import (
	"fmt"
)

/*
You are given two linked lists representing two non-negative numbers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

func main() {
	input1 := []int{2, 4, 3}
	input2 := []int{5, 6, 4}

	fmt.Printf("%v\n", add(input1, input2))

}

func add(a, b []int) []int {
	carry := 0
	var result []int

	for i := 0; i < len(a) || i < len(b); i++ {
		ax, bx := 0, 0
		if i < len(a) {
			ax = a[i]
		}
		if i < len(b) {
			bx = b[i]
		}
		r := ax + bx + carry
		newcarry := 0
		for r >= 10 {
			r = r - 10
			newcarry += 1
		}
		result = append(result, r)
		carry = newcarry
	}
	if carry > 0 {
		result[len(result)-1] = result[len(result)-1] + carry
	}
	return result
}
