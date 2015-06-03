package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var test []int
	for i := 0; i < 1000; i++ {
		test = append(test, rand.Int())
	}
	fmt.Printf("%v", top100(test))
}

func top100(universe []int) []int {
	var result [100]int
	for i := range universe {
		prev := -1
		for j := len(result) - 1; j >= 0; j-- {
			if result[j] < universe[i] && prev == -1 {
				result[j] = universe[i]
				prev = result[j]
			} else if prev == -1 && result[j] > universe[i] {
				continue
			} else if prev > result[j] {
				newPrev := result[j]
				result[j] = prev
				prev = newPrev
			} else {
				break
			}
		}
	}
	return result[:]
}
