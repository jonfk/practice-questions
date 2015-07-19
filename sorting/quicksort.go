package mysorting

import (
	//"fmt"
	"math/rand"
)

// func main() {
// 	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
// 	b := []int{1, 7, 6, 9, 5, 89, 88, 66, 54, 437, 8100}
// 	fmt.Printf("%v\n", qsort(a))
// 	fmt.Printf("%v\n", qsort(b))
// }

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	lo, hi := 0, len(a)-1

	// choose pivot
	pivot := rand.Int() % (len(a))

	a[pivot], a[hi] = a[hi], a[pivot]

	for i := range a {
		if a[i] < a[hi] {
			a[i], a[lo] = a[lo], a[i]
			lo++
		}
	}
	a[hi], a[lo] = a[lo], a[hi]
	qsort(a[:lo])
	qsort(a[lo:])
	return a
}
