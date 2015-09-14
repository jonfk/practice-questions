package main

import (
	"fmt"
	"math/rand"
)

type Person struct {
	H int
	W int
}

func qsort(a []Person) []Person {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]
	for i := 0; i < right; i++ {
		if a[i].H < a[right].H {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[right], a[left] = a[left], a[right]
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func main() {
	test := []Person{Person{65, 100}, Person{70, 150}, Person{56, 90}, Person{75, 190}, Person{60, 95}, Person{68, 110}, Person{57, 120}}
	fmt.Printf("%v\n", qsort(test))
	fmt.Printf("max tower %v\n", tower(test))
}

func tower(list []Person) []Person {
	sortedByHeight := qsort(list)

	var potentialTowers [][]Person

	for i := range sortedByHeight {
		canditate := sortedByHeight[i]
		for j := range potentialTowers {
			if last(potentialTowers[j]).W < canditate.W {
				save := make([]Person, len(potentialTowers[j]), cap(potentialTowers[j]))
				copy(save, potentialTowers[j])
				potentialTowers = append(potentialTowers, save)
				potentialTowers[j] = append(potentialTowers[j], canditate)
			}
		}
		potentialTowers = append(potentialTowers, []Person{canditate})
	}
	fmt.Printf("potential %v\n", potentialTowers)
	maxHeight := 0
	var maxStack []Person
	for i := range potentialTowers {
		if len(potentialTowers[i]) > maxHeight {
			maxHeight = len(potentialTowers[i])
			maxStack = potentialTowers[i]
		}
	}
	return maxStack
}

func last(a []Person) Person {
	return a[len(a)-1]
}
