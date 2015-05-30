package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}}
	fmt.Printf("%v\n", matrix)
	rotate(matrix)
	fmt.Printf("%v\n", matrix)

}

func rotate(matrix [][]int) {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first

			//save top
			top := matrix[first][i]

			matrix[first][i] = matrix[last-offset][first]

			matrix[last-offset][first] = matrix[last][last-offset]

			matrix[last][last-offset] = matrix[i][last]

			matrix[i][last] = top
		}
	}
}
