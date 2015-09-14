package main

import (
	"fmt"
)

type Sub struct {
	Start int
	End   int
}

func toSort(a []int) (int, int) {
	start := 0
	last := a[0]
	var subseqs []Sub

	for i := range a {
		if a[i] >= last {
			last = a[i]
		} else {
			subseqs = append(subseqs, Sub{start, i - 1})
			start = i
			last = a[i]
		}
	}
	if subseqs[len(subseqs)-1].Start != start {
		subseqs = append(subseqs, Sub{start, len(a) - 1})
	}
	fmt.Printf("%v\n", subseqs)

	if len(subseqs) < 2 {
		return 0, 0
	} else if len(subseqs) == 2 {
		return 0, len(a) - 1
	} else {
		start := minStart(subseqs[1 : len(subseqs)-1])
		end := maxEnd(subseqs[1 : len(subseqs)-1])
		var i int
		for i = range a {
			if a[i] > start {
				break
			}
		}
		finalStart := i

		for i = len(a) - 1; i >= 0; i-- {
			if a[i] < end {
				break
			}
		}
		finalEnd := i + 1
		return finalStart, finalEnd
	}
}

func minStart(a []Sub) int {
	if len(a) < 1 {
		return -1
	}
	min := a[0].Start
	for i := range a {
		if a[i].Start < min {
			min = a[i].Start
		}
	}
	return min
}

func maxEnd(a []Sub) int {
	if len(a) < 1 {
		return -1
	}
	max := a[0].End
	for i := range a {
		if a[i].End > max {
			max = a[i].End
		}
	}
	return max
}

func main() {
	test := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	a, b := toSort(test)

	fmt.Printf("%d %d\n", a, b)
}
