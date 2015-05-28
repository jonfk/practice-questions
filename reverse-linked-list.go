package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type LinkList struct {
	Value int
	Next  *LinkList
}

func (a *LinkList) Append(x int) *LinkList {
	p := a
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &LinkList{Value: x}
	return a
}

func main() {
	fmt.Printf("reverse-linklist\n")

	list := toLinkList([]int{1, 2, 3, 4, 5})

	spew.Printf("%#v\n", list)
	spew.Printf("%#v\n", reverse(list))
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

func reverse(a *LinkList) *LinkList {
	if a.Next == nil {
		return a
	}
	reversedTail := reverse(a.Next)
	result := reversedTail.Append(a.Value)
	return result
}
