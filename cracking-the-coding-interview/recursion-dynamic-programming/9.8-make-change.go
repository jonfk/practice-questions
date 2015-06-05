package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", makeChange(200, []int{200, 100, 50, 20, 10, 5, 2, 1}))
}

func makeChange(amount int, denoms []int) int {
	if len(denoms) <= 1 {
		return 1
	} else {
		var ways int
		denomAmount := denoms[0]
		for i := 0; i*denomAmount <= amount; i++ {
			remainingAmount := amount - (denomAmount * i)
			ways += makeChange(remainingAmount, denoms[1:])
		}
		return ways
	}
}
