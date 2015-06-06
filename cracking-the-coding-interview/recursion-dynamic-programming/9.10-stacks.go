package main

import (
	"fmt"
	"math/rand"
)

type Box struct {
	W int
	L int
	H int
}

func createStack(boxes []Box, bottom Box) (int, []Box) {

	// if len(boxes) < 2 {
	// 	return 2, append([]Box{bottom}, boxes...)
	// }

	var (
		maxHeight int
		maxStack  []Box
	)

	for i := range boxes {
		box := boxes[i]
		if canBeAbove(boxes[i], bottom) {
			remainingBoxes := append(boxes[:i], boxes[i+1:]...)
			stackHeight, newStack := createStack(remainingBoxes, box)
			if stackHeight > maxHeight {
				maxHeight = stackHeight
				maxStack = newStack
			}
		}
	}

	if bottom.H == 0 && bottom.L == 0 && bottom.W == 0 {
		return len(maxStack), maxStack
	}
	maxStack = append([]Box{bottom}, maxStack...)
	return len(maxStack), maxStack
}

func canBeAbove(above, bottom Box) bool {
	if above.W > bottom.W && above.L > bottom.L && above.H > bottom.H {
		return true
	}
	return false
}

func main() {
	var test []Box
	for i := 0; i < 1000; i++ {
		h, w, l := rand.Int()%100, rand.Int()%100, rand.Int()%100
		box := Box{w + i, l + i, h + i}
		test = append(test, box)
	}
	fmt.Printf("Boxes: %v\n", test)

	var (
		maxHeight int
		maxStack  []Box
	)
	for i := range test {
		remainingBoxes := append(test[:i], test[i+1:]...)
		stackHeight, newStack := createStack(remainingBoxes, test[i])
		if stackHeight > maxHeight {
			maxHeight = stackHeight
			maxStack = newStack
		}
	}
	fmt.Printf("maxStack: %v\nheight: %d\n", maxStack, maxHeight)
}
