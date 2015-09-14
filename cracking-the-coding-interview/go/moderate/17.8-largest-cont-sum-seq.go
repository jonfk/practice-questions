package main

import "fmt"

func maxSum(a []int) int {
	sum, maxSum := 0, 0
	for i := range a {
		sum += a[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func main() {
	test := []int{2, -8, 3, -2, 4, -10}
	fmt.Printf("%d\n", maxSum(test))
}
