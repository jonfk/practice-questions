package main

import (
	"fmt"
)

/*
 Recently I attended the interview at Google and I was asked
"You are given a sorted list of disjoint intervals
and an interval, e.g. [(1, 5), (10, 15), (20, 25)] and (12, 27).
Your task is to merge them into a sorted list of disjoint intervals: [(1, 5), (10, 27)]."
*/

type Pair struct {
	A, B int
}

func main() {
	disjoint := []Pair{Pair{1, 5}, Pair{10, 15}, Pair{20, 25}}
	fmt.Printf("%v\n", disjoint)
	fmt.Printf("%v\n", merge(disjoint, Pair{12, 27}))
}

func merge(disjoint []Pair, interval Pair) []Pair {
	for i := range disjoint {
		if interval.A > disjoint[i].A && interval.A < disjoint[i].B {
			disjoint[i] = Pair{disjoint[i].A, interval.B}
			break
		}
	}
	for i := range disjoint {
		if i < len(disjoint)-1 && disjoint[i].B > disjoint[i+1].A {
			disjoint[i] = Pair{disjoint[i].A, max(disjoint[i+1].B, disjoint[i].B)}
			disjoint = append(disjoint[:i+1], disjoint[i+2:]...)
		}
	}
	return disjoint
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
