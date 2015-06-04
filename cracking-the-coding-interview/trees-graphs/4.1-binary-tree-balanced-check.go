package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	//"log"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func sortedArrayToBST(a []int) *Tree {
	if len(a) == 0 {
		return nil
	}
	if len(a) < 2 {
		return &Tree{Left: nil, Value: a[0], Right: nil}
	}
	mid := len(a) / 2
	left := sortedArrayToBST(a[:mid])
	right := sortedArrayToBST(a[mid+1:])
	return &Tree{Left: left, Value: a[mid], Right: right}
}

func main() {
	fmt.Printf("balance-tree\n")
	balancedTree := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	// unbalancedTree := &Tree{
	// 	Left: &Tree{
	// 		Value: 1,
	// 	},
	// 	Value: 2,
	// 	Right: &Tree{
	// 		Left: &Tree{
	// 			Left:  &Tree{Value: 5},
	// 			Value: 6,
	// 		},
	// 		Value: 7,
	// 		Right: &Tree{
	// 			Value: 8,
	// 			Right: &Tree{
	// 				Value: 9,
	// 				Right: &Tree{Value: 10}}},
	// 	},
	// }

	unbalancedTree2 := &Tree{
		Left: &Tree{
			Value: 1,
			Right: &Tree{
				Value: 8,
				Right: &Tree{
					Value: 9,
					Right: &Tree{Value: 10}}},
		},
		Value: 2,
		Right: &Tree{
			Left: &Tree{
				Left:  &Tree{Value: 5},
				Value: 6,
			},
			Value: 7,
			Right: &Tree{
				Value: 8,
				Right: &Tree{
					Value: 9,
					Right: &Tree{Value: 10}}},
		},
	}
	fmt.Printf("true: %v\nfalse: %v\n", checkBalanced(balancedTree), checkBalanced(unbalancedTree2))
}

func checkBalanced(a *Tree) bool {
	if a == nil {
		return true
	}
	if abs(maxDepth(a.Left)-maxDepth(a.Right)) < 2 {
		return checkBalanced(a.Left) && checkBalanced(a.Right)
	} else {
		return false
	}
}

func maxDepth(a *Tree) int {
	if a == nil {
		return 0
	} else {
		return 1 + max(maxDepth(a.Left), maxDepth(a.Right))
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
