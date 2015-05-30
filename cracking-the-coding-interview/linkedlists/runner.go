package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type LinkList struct {
	Value int
	Next  *LinkList
}

func toLinkList(a []int) *LinkList {
	var result, p *LinkList

	for i := range a {
		if p == nil {
			p = &LinkList{Value: a[i]}
			result = p
		} else {
			p.Next = &LinkList{Value: a[i]}
			p = p.Next
		}

	}
	return result
}

func main() {
	fmt.Printf("linklist-runner\n")
	test := toLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	test2 := toLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	spew.Printf("%v\n", test)
	fmt.Printf("%v, %v \n", median(test), median(test2))
}

func median(a *LinkList) int {
	if a == nil {
		return -1
	}
	var p, q *LinkList

	p, q = a, a

	for i := 0; p != nil; i++ {
		p = p.Next
		if i == 1 {
			q = q.Next
			i = -1
		}
	}
	return q.Value
}
