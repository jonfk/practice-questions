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
	fmt.Printf("circular-list\n")
	test := &LinkList{Value: 1, Next: &LinkList{Value: 2, Next: &LinkList{Value: 3}}}
	test2 := &LinkList{Value: 4, Next: &LinkList{Value: 5, Next: &LinkList{Value: 6}}}
	test.Next.Next.Next = test2
	test2.Next.Next.Next = test2
	spew.Printf("%v\n", test)
	spew.Printf("%v\n", circular(test))
}

func circular(a *LinkList) *LinkList {
	var set map[*LinkList]int = make(map[*LinkList]int)
	p := a
	for p != nil {
		set[p] = set[p] + 1
		if set[p] > 1 {
			return p
		}
		p = p.Next
	}
	return nil
}
