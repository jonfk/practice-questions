package mysorting

import (
//"fmt"
)

// func main() {
// 	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
// 	//array := []int{2, 1}
// 	fmt.Printf("%v\n", array)
// 	fmt.Printf("%v\n", mergesort(array))
// }

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
