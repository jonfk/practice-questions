package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

type Stack struct {
	Value []int
}

func (a *Stack) Pop() (int, error) {
	if len(a.Value) < 1 {
		return -1, fmt.Errorf("Pop() on empty Stack")
	}
	var x int
	x, a.Value = a.Value[len(a.Value)-1], a.Value[:len(a.Value)-1]
	return x, nil
}

func (a *Stack) Peek() (int, error) {
	if a == nil || len(a.Value) < 1 {
		return -1, fmt.Errorf("Peek on nil stack")
	}
	return a.Value[len(a.Value)-1], nil
}

func (a *Stack) Push(x int) error {
	if a == nil {
		return fmt.Errorf("Push on nil stack")
	}
	q, err := a.Peek()
	if err != nil {
		q = 1000
	}
	if q < x {
		return fmt.Errorf("violating rules of towers of hanoi")
	}
	a.Value = append(a.Value, x)
	return nil
}

func main() {
	tower1, tower2, tower3 := new(Stack), new(Stack), new(Stack)
	err := tower1.Push(3)
	if err != nil {
		log.Fatal(err)
	}
	err = tower1.Push(2)
	if err != nil {
		log.Fatal(err)
	}
	err = tower1.Push(1)
	if err != nil {
		log.Fatal(err)
	}
	moveTower(3, tower1, tower3, tower2)
	spew.Printf("%v %v %v\n", tower1, tower2, tower3)
}

func moveTower(numDisk int, src, dst, inter *Stack) {
	spew.Printf("%v %v %v\n", src, dst, inter)
	if numDisk == 1 {
		x, err := src.Pop()
		if err != nil {
			log.Fatal(err)
		}
		err = dst.Push(x)
		if err != nil {
			log.Fatal(err)
		}
		return
	} else {
		moveTower(numDisk-1, src, inter, dst)
		moveTower(1, src, dst, inter)
		moveTower(numDisk-1, inter, dst, src)
		return
	}
}
