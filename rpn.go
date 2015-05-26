package main

/*

Evaluate the value of an arithmetic expression in Reverse Polish Notation. Valid operators are +, -, *, /. Each operand may be an integer or another expression. For example:

  ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
  ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6

*/

import (
	"fmt"
	"strconv"
)

type Stack struct {
	V []int
}

func (s *Stack) Push(a int) {
	s.V = append(s.V, a)
}
func (s *Stack) Pop() int {
	var x int
	x, s.V = s.V[len(s.V)-1], s.V[:len(s.V)-1]
	return x
}

func main() {
	input := []string{"4", "13", "5", "/", "+"}
	input2 := []string{"2", "1", "+", "3", "*"}
	fmt.Printf("%v\n", rpn(input))
	fmt.Printf("%v\n", rpn(input2))
}

func rpn(input []string) int {
	stack := &Stack{}
	for i := range input {
		switch input[i] {
		case "*":
			x := stack.Pop()
			y := stack.Pop()
			stack.Push(y * x)
		case "+":
			x := stack.Pop()
			y := stack.Pop()
			stack.Push(y + x)
		case "/":
			x := stack.Pop()
			y := stack.Pop()
			stack.Push(y / x)
		default:
			x, err := strconv.Atoi(input[i])
			if err != nil {
				fmt.Printf("%v", err)
			}
			stack.Push(x)
		}
	}
	return stack.Pop()

}
