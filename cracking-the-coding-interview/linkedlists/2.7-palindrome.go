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
	fmt.Printf("linklist-palindrome\n")
	test := toLinkList([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})
	test2 := toLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	spew.Printf("%v\n", test)
	fmt.Printf("true: %v\nfalse: %v\n", palindrome(test), palindrome(test2))
}

func (a *LinkList) equal(b *LinkList) bool {
	p, q := a, b
	for p != nil && q != nil {
		if p.Value != q.Value {
			return false
		}
		p = p.Next
		q = q.Next
	}
	if p != nil || q != nil {
		return false
	}
	return true
}

func reverse(a *LinkList) *LinkList {
	var result *LinkList
	for a != nil {
		result = &LinkList{Value: a.Value, Next: result}
		a = a.Next
	}
	return result
}

func palindrome(a *LinkList) bool {
	reverse := reverse(a)
	return reverse.equal(a)
}
