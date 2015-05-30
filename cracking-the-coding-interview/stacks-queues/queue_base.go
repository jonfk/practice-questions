package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

// LinkList style
/*
type Queue struct {
	First, Last *node
}

type node struct {
	Value int
	next  *node
}

func (a *Queue) Enqueue(x int) error {
	if a == nil {
		return fmt.Errorf("enqueuing on nil queue")
	}
	newNode := &node{Value: x}
	if a.First == nil && a.Last == nil {
		a.First = newNode
		a.Last = newNode
	} else {
		a.Last.next = newNode
		a.Last = newNode
	}
	return nil
}

func (a *Queue) Dequeue() (int, error) {
	if a == nil {
		return -1, fmt.Errorf("dequeuing on nil queue")
	}
	val := a.First.Value
	a.First = a.First.next
	return val, nil
}
*/

// slice style
///*
type Queue struct {
	Value []int
}

func (a *Queue) Enqueue(x int) error {
	if a == nil {
		return fmt.Errorf("enqueuing on nil queue")
	}
	a.Value = append(a.Value, x)
	return nil
}
func (a *Queue) Dequeue() (int, error) {
	if a == nil || len(a.Value) < 1 {
		return -1, fmt.Errorf("Dequeuing on nil or empty queue")
	}
	var x int
	x, a.Value = a.Value[0], a.Value[1:]
	return x, nil
}

//*/

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	spew.Printf("%v\n", queue)
	queue.Dequeue()
	x, _ := queue.Dequeue()
	fmt.Printf("%d\n", x)
}
