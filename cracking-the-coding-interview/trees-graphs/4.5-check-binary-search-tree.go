package main

import (
	"fmt"
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
	fmt.Printf("check-binary-tree-bst\n")
	bst := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	nonbst := sortedArrayToBST([]int{10, 9, 7, 5, 8, 6, 2, 6, 9})
	fmt.Printf("true: %v\nfalse: %v\n", IsBinarySearchTree(bst), IsBinarySearchTree(nonbst))
}

func IsBinarySearchTree(a *Tree) bool {
	if a == nil {
		return true
	} else if a.Left == nil && a.Right == nil {
		return true
	} else if a.Left != nil && a.Left.Value > a.Value {
		return false
	} else if a.Right != nil && a.Right.Value < a.Value {
		return false
	} else {
		return IsBinarySearchTree(a.Left) && IsBinarySearchTree(a.Right)
	}
}
