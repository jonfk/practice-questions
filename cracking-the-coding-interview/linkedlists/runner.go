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
	//fmt.Printf("%v, %v \n", median(test), median(test2))
	fmt.Printf("%v, %v \n", median2(test), median2(test2))
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

func median2(a *LinkList) int {
	if a == nil {
		return -1
	}
	p, q := a, a
	if p.Next.Next == nil {
		return p.Value
	}

	for p != nil {
		if p.Next != nil {
			p = p.Next.Next
		} else {
			p = nil
			break
		}
		q = q.Next
	}
	return q.Value
}
