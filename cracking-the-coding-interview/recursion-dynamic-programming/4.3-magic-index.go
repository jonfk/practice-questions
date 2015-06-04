package main

import (
	"fmt"
	"log"
)

func main() {
	test := []int{-1, -2, -2, -3, 3, 4, 5, 6, 7, 8, 10, 20, 54}
	m, err := magic(test, 0, len(test)-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", m)

}

func magicNaive(n []int) (int, error) {
	for i := range n {
		if n[i] == i {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Not found")
}

func magic(n []int, start, end int) (int, error) {
	if end < start || start < 0 || end >= len(n) {
		return -1, fmt.Errorf("Not found")
	}

	mid := (start + end) / 2
	if n[mid] == mid {
		return mid, nil
	} else {
		if n[mid] < mid {
			return magic(n, mid+1, end)
		} else {
			return magic(n, start, mid-1)
		}
	}
}
