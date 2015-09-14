package main

/*
You are given two linked lists representing two non-negative numbers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/
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
	test := toLinkList([]int{7, 1, 6, 9})
	test2 := toLinkList([]int{5, 9, 2})
	spew.Printf("%v\n", add(test, test2))
}

func add(a, b *LinkList) *LinkList {
	p, q := a, b
	carry := 0
	var result, z *LinkList
	for p != nil && q != nil {
		temp := p.Value + q.Value + carry
		carry = 0
		for temp >= 10 {
			temp -= 10
			carry += 1
		}
		if result == nil {
			result = &LinkList{Value: temp}
			z = result
		} else {
			z.Next = &LinkList{Value: temp}
			z = z.Next
		}
		p = p.Next
		q = q.Next
	}
	var r *LinkList
	if p != nil {
		r = p
	} else {
		r = q
	}
	p1 := r
	for carry != 0 {
		if p1 == nil {
			r = &LinkList{Value: carry}
			carry = 0
			continue
		}
		temp := p1.Value + carry
		for temp > 10 {
			temp -= 10
			carry += 1
		}
		p1.Value = temp
		p1 = p1.Next
	}
	for r != nil {
		z.Next = &LinkList{Value: r.Value}
		z = z.Next
		r = r.Next
	}
	return result
}
