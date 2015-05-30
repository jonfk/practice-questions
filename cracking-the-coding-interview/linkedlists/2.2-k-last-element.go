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
	fmt.Printf("k-last-element\n")
	test := toLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	spew.Printf("%v\n", test)
	fmt.Printf("%d\n", klast(test, 3))
}

func klast(a *LinkList, k int) int {
	p := a
	for i := 0; i < k; i++ {
		p = p.Next
	}
	q := a
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q.Value
}
