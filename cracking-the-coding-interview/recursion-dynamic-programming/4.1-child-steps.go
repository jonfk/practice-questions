package main

import (
	"fmt"
)

type Value struct {
	Value int
}

var ways map[int]*Value

func init() {
	ways = make(map[int]*Value)
}

func main() {
	fmt.Printf("%d\n", waysSteps(100))
	//fmt.Printf("%d\n", naive(100))
}

func waysSteps(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	} else {
		if ways[n] != nil {
			return ways[n].Value
		} else {
			way := waysSteps(n-1) + waysSteps(n-2) + waysSteps(n-3)
			ways[n] = &Value{Value: way}
			return way
		}
	}
}

func naive(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	} else {
		return waysSteps(n-1) + waysSteps(n-2) + waysSteps(n-3)
	}
}
