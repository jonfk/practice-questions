package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func InitStack() *Stack {
	return &Stack{}
}

// LinkList style stack
/*
type Stack struct {
	Top *node
}
type node struct {
	Value int
	Next  *node
}

func (a *Stack) Pop() (int, error) {
	if a == nil || a.Top == nil {
		return -1, fmt.Errorf("Pop() on empty Stack")
	}
	re := a.Top.Value
	a.Top = a.Top.Next
	return re, nil
}

func (a *Stack) Push(x int) {
	if a == nil {
		a = &Stack{}
	}
	ntop := &node{Value: x, Next: a.Top}
	a.Top = ntop
}

//*/

///*
// slice style stack

type Stack struct {
	Value []int
}

func (a *Stack) Pop() (int, error) {
	if len(a.Value) < 1 {
		return -1, fmt.Errorf("Pop() on empty Stack")
	}
	var x int
	x, a.Value = a.Value[len(a.Value)-1], a.Value[:len(a.Value)-1]
	return x, nil
}

func (a *Stack) Push(x int) {
	a.Value = append(a.Value, x)
}

//*/

func main() {
	var stack *Stack = &Stack{}
	stack.Push(1)
	spew.Printf("%v\n", stack)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	a, _ := stack.Pop()
	fmt.Printf("%d\n", a)
}
