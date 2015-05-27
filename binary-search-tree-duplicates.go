package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

/*
Given a binary search tree where there may contain
duplicates, but all other logic of the BST is intact,
determine the most frequently occurring element.
*/

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

type Stack struct {
	Value []int
}

func (s *Stack) Pop() int {
	var x int
	x, s.Value = s.Value[len(s.Value)-1], s.Value[:len(s.Value)-1]
	return x
}

func (s *Stack) Push(x int) {
	s.Value = append(s.Value, x)
}

type TreeStack struct {
	Value []*Tree
}

func (s *TreeStack) Pop() *Tree {
	var x *Tree
	x, s.Value = s.Value[len(s.Value)-1], s.Value[:len(s.Value)-1]
	return x
}

func (s *TreeStack) Push(x *Tree) {
	s.Value = append(s.Value, x)
}

type TreeQueue struct {
	Value []*Tree
}

func (q *TreeQueue) enqueue(x *Tree) {
	q.Value = append(q.Value, x)
}

func (q *TreeQueue) dequeue() *Tree {
	var x *Tree
	x, q.Value = q.Value[0], q.Value[1:]
	return x
}

func main() {
	/*
	      4
	    2    6
	   1 3  5 6
	*/
	bst := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 12, 12, 12, 13, 14})
	spew.Printf("%#v\n", bst)

	fmt.Printf("%v\n", mostFrequent(bst))
}

func mostFrequent(a *Tree) int {
	var freq map[int]int = make(map[int]int)
	stack := new(TreeStack)
	stack.Push(a)
	for len(stack.Value) > 0 {
		popped := stack.Pop()

		if popped != nil {
			freq[popped.Value] = freq[popped.Value] + 1
			stack.Push(popped.Right)
			stack.Push(popped.Left)
		}
	}
	maxKey, maxValue := 0, 0
	for key, value := range freq {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}
	return maxKey
}

func DFS(a *Tree) []int {
	if a == nil {
		return nil
	}
	result := []int{a.Value}
	left := DFS(a.Left)
	right := DFS(a.Right)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func DFSStack(a *Tree) []int {
	var result []int
	stack := new(TreeStack)
	stack.Push(a)
	for len(stack.Value) > 0 {
		popped := stack.Pop()

		if popped != nil {
			result = append(result, popped.Value)
			stack.Push(popped.Right)
			stack.Push(popped.Left)
		}
	}
	return result
}

// left child -> parent -> right child
func inOrderTraversal(a *Tree) []int {
	var result []int
	stack := new(TreeStack)

	p := a
	for len(stack.Value) > 0 || p != nil {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			popped := stack.Pop()
			result = append(result, popped.Value)
			p = popped.Right
		}

	}
	return result
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
