package main

import (
	"fmt"
)

func main() {

	test := []int{4, 3, 2, 1} // 6
	fmt.Printf("%d\n", inversions1(test))
	sorted, counted := count(test)
	fmt.Printf("%v %d\n", sorted, counted)
	fmt.Printf("%d\n", mergeAndCountSplit([]int{3, 4}, []int{1, 2}))
}

func inversions1(a []int) int {
	result := 0
	for i := range a {
		x := a[i]
		for j := range a {
			if x > a[j] {
				result++
			}
		}
	}
	return result
}

// divide and conquer
func count(a []int) ([]int, int) {
	if len(a) < 2 {
		return a, 0
	}
	mid := len(a) / 2
	xSorted, x := count(a[:mid])
	ySorted, y := count(a[mid:])
	z := mergeAndCountSplit(xSorted, ySorted)
	return mergesort(a), (x + y + z)
}

// incorrect
func mergeAndCountSplit(a, b []int) int {
	i, j, result := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			result++
			j++
		} else {
			if j == len(b)-1 {
				i++
				j = 0
			} else {
				j++
			}
		}
	}
	return result
}

func mergesort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2

	first := mergesort(array[:mid])
	second := mergesort(array[mid:])
	return merge(first, second)
}

func merge(arr1, arr2 []int) []int {
	var result []int

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	// at least one of arr1 or arr2 must be exhausted
	if i < len(arr1) {
		for ; i < len(arr1); i++ {
			result = append(result, arr1[i])
		}
	}
	if j < len(arr2) {
		for ; i < len(arr2); j++ {
			result = append(result, arr2[j])
		}
	}
	return result

}
