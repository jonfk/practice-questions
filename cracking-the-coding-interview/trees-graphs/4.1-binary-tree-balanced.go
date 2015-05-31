package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
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

func balance(a *Tree) {
	leftDepth, rightDepth := maxDepth(a.Left), maxDepth(a.Right)
	if abs(leftDepth-rightDepth) < 2 {
		return
	}
	if leftDepth < rightDepth {
		leaf := removeDeepLeaf(a.Right)
		a.Left.add(leaf)
	} else {
		leaf := removeDeepLeaf(a.Left)
		a.Right.add(leaf)
	}
	balance(a)

}

func (a *Tree) add(x *Tree) {
	if a.Left == nil {
		a.Left = x
	} else if a.Right == nil {
		a.Right = x
	} else {
		leftDepth, rightDepth := maxDepth(a.Left), maxDepth(a.Right)
		if leftDepth > rightDepth {
			a.Right.add(x)
		} else {
			a.Left.add(x)
		}
	}
}

func getDeepLeaf(a *Tree) *Tree {
	if a.Left == nil && a.Right == nil {
		leaf := a
		a = nil
		return leaf
	} else {
		leftDepth, rightDepth := maxDepth(a.Left), maxDepth(a.Right)
		if leftDepth >= rightDepth {
			return getDeepLeaf(a.Left)
		} else {
			return getDeepLeaf(a.Right)
		}
	}
}

func removeNode(a, x *Tree) bool {
	if a == nil {
		return false
	} else if a.Left == x {
		a.Left = nil
		return true
	} else if a.Right == x {
		a.Right = nil
		return true
	} else {
		return removeNode(a.Left, x) || removeNode(a.Right, x)
	}
}

func removeDeepLeaf(a *Tree) *Tree {
	leaf := getDeepLeaf(a)
	if !removeNode(a, leaf) {
		log.Fatal("Leaf not found!")
	}
	return leaf
}

func main() {
	fmt.Printf("balance-tree")
	balancedTree := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	unbalancedTree := &Tree{
		Left: &Tree{
			Value: 1,
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

	//spew.Printf("balanced: %#v\n\nUnbalanced: %v\n\n", balancedTree, unbalancedTree)
	fmt.Printf(" before Balanced: %s\n", spew.Sdump(balancedTree))
	//x := removeDeepLeaf(unbalancedTree)
	//fmt.Printf("Deep Leaf : %v", spew.Sdump(x))
	fmt.Printf("UnBalanced: %s\n", spew.Sdump(unbalancedTree))
	fmt.Printf("maxDepth: %d\n", maxDepth(unbalancedTree))

	balance(unbalancedTree)
	fmt.Printf("after unbalanced: %s\n", spew.Sdump(unbalancedTree))

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
