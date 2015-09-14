package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Stack struct {
	Top *node
}
type node struct {
	Value int
	min   int
	Next  *node
}

func (a *Stack) Min() (int, error) {
	if a == nil || a.Top == nil {
		return -1, fmt.Errorf("Min() on empty Stack")
	}
	return a.Top.min, nil
}

func (a *Stack) Pop() (int, error) {
	if a == nil || a.Top == nil {
		return -1, fmt.Errorf("Pop() on empty Stack")
	}
	re := a.Top.Value
	a.Top = a.Top.Next
	return re, nil
}

func (a *Stack) Push(x int) error {
	if a == nil {
		return fmt.Errorf("pushing on nil stack")
	}
	var ntop *node
	if a.Top == nil || a.Top.min > x {
		ntop = &node{Value: x, min: x, Next: a.Top}
	} else {
		ntop = &node{Value: x, min: a.Top.min, Next: a.Top}
	}
	a.Top = ntop
	return nil
}

func main() {
	var stack *Stack = new(Stack)
	stack.Push(2)
	spew.Printf("%v\n", stack)
	stack.Push(2)
	stack.Push(3)
	min, _ := stack.Min()
	fmt.Printf("min 2: %d\n", min)
	stack.Push(1)
	min, _ = stack.Min()
	fmt.Printf("min 1: %d\n", min)
	stack.Pop()
	min, _ = stack.Min()
	fmt.Printf("min 2: %d\n", min)
}
