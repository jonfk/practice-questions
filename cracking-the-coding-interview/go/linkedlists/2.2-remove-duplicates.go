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

func remDup(a *LinkList) *LinkList {
	var set map[int]int = make(map[int]int)
	p := a
	var q *LinkList
	for p != nil {
		set[p.Value] = set[p.Value] + 1
		if set[p.Value] > 1 {
			if q == nil {
				a = a.Next
				p = a
				continue
			} else {
				q.Next = p.Next
				p = p.Next
				continue
			}
			set[p.Value] = set[p.Value] - 1
		}
		q = p
		p = p.Next
	}
	return a
}

func main() {
	fmt.Printf("linklist-runner\n")
	test := toLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	test2 := toLinkList([]int{1, 2, 2, 2, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8, 9, 10, 10})
	spew.Printf("%v\n", remDup(test))
	spew.Printf("%v\n", remDup(test2))

}
