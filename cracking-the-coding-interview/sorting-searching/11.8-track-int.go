package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (a *Tree) Insert(x int) *Tree {
	if a == nil {
		return &Tree{Value: x}
	} else if a.Value >= x {
		a.Left = a.Left.Insert(x)
	} else {
		a.Right = a.Right.Insert(x)
	}
	return a
}

func main() {
	a := &Ranks{}
	a.track(5)
	a.track(1)
	a.track(4)
	a.track(4)
	a.track(5)
	a.track(9)
	a.track(7)
	a.track(13)
	a.track(3)
	spew.Dump(a)

	fmt.Printf("Rank: %d\n", a.getRank(4))

}

type Ranks struct {
	List *Tree
}

func (a *Ranks) track(x int) {
	if a.List == nil {
		a.List = a.List.Insert(x)
	} else {
		a.List = a.List.Insert(x)

	}
}

func (a *Ranks) getRank(x int) int {
	return getRankTree(a.List, x)
}

func getRankTree(t *Tree, x int) int {
	if t == nil {
		return 0
	}
	if t.Value == x {
		return size(t.Left)
	} else if t.Value > x {
		return getRankTree(t.Left, x)
	} else {
		return size(t.Left) + 1 + getRankTree(t.Right, x)
	}
}

func size(t *Tree) int {
	if t == nil {
		return 0
	} else {
		return size(t.Left) + 1 + size(t.Right)
	}
}
