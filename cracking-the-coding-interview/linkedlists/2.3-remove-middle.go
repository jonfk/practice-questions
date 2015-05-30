package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

/*
Implement an algorithm to delete a node in the
middle of a singly linked list, given only access to that node.
*/

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
	fmt.Printf("linklist-delete-middle\n")
	test := toLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	spew.Printf("%v\n", test)
	a := median(test)
	spew.Printf("%v \n", a)
	deleteMiddle(a)
	spew.Printf("%v \n", test)
}

func deleteMiddle(a *LinkList) {
	p := a
	for p.Next != nil {
		p.Value = p.Next.Value
		p = p.Next
	}
}

func median(a *LinkList) *LinkList {
	if a == nil {
		return nil
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
	return q
}
