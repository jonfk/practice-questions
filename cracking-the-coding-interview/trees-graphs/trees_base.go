package main

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
