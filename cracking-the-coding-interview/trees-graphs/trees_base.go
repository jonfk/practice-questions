package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

/*
Issues to Watch out for:

- Binary Tree vs Binary search tree
- balanced vs unbalanced: Balanced means small difference in depth in subtrees
- Full and complete: all leaves are at the bottom of tree. tree must have exactly 2^n-1 nodes

Binary tree traversal
- in-order traversal
- post-order traversal
- pre-order traversal

Tree balancing: Red-black trees and avl trees

Tries: n-ary trees where each path down the tree represents a word


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

func postOrderTraversal(a *Tree) []int {
	if a == nil {
		return nil
	}
	left := postOrderTraversal(a.Left)
	right := postOrderTraversal(a.Right)
	temp := append(left, right...)
	return append(temp, a.Value)
}
func inOrderTraversal(a *Tree) []int {
	if a == nil {
		return nil
	}
	left := postOrderTraversal(a.Left)
	temp := append(left, a.Value)
	right := postOrderTraversal(a.Right)
	return append(temp, right...)
}
func preOrderTraversal(a *Tree) []int {
	if a == nil {
		return nil
	}
	left := postOrderTraversal(a.Left)
	temp := append([]int{a.Value}, left...)
	right := postOrderTraversal(a.Right)
	return append(temp, right...)
}

func main() {
	tree := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	spew.Dump(tree)
	fmt.Printf("post: %v\n", postOrderTraversal(tree))
	fmt.Printf("pre: %v\n", preOrderTraversal(tree))
	fmt.Printf("in: %v\n", inOrderTraversal(tree))
}
