package main

import (
	"fmt"
)

/*
Given a binary tree, design an algorithm which creates a linked list of all the
nodes at each depth (e.g., if you have a tree with depth D, you'll have D linked
lists).
*/

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

func nodesAtDepth(a *Tree, depth int) []int {
	if a == nil && depth > 1 {
		return []int{}
	} else if depth == 1 {
		return []int{a.Value}
	} else {
		result := append(nodesAtDepth(a.Left, depth-1), nodesAtDepth(a.Right, depth-1)...)
		return result
	}
}

func main() {
	tree := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	fmt.Printf("%v\n", nodesAtDepth(tree, 2))
}
