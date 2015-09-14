package main

import (
	"fmt"
)

type Slot int

const (
	R = iota
	G
	B
	Y
)

var actual []Slot = []Slot{R, G, B, Y}

func checkMasterMind(guess []Slot) (int, int) {
	copyActual := make([]Slot, len(actual))
	copy(copyActual, actual)
	hits, pseudoHits := 0, 0

	toDeleteGuess := []int{}
	toDeleteActual := []int{}
	for i := range guess {
		if guess[i] == copyActual[i] {
			hits++

			toDeleteGuess = append(toDeleteGuess, i)
			toDeleteActual = append(toDeleteActual, i)
		}
	}
	for i := range toDeleteGuess {
		index := toDeleteGuess[i]
		guess = append(guess[:index], guess[index+1:]...)
	}
	for i := range toDeleteActual {
		index := toDeleteActual[i]
		copyActual = append(copyActual[:index], copyActual[index+1:]...)
	}

	for i := range guess {
		index, found := find(copyActual, guess[i])
		if found {
			copyActual = append(copyActual[:index], copyActual[index+1:]...)
			pseudoHits++
		}
	}
	return hits, pseudoHits
}

func find(l []Slot, x Slot) (int, bool) {
	for i := range l {
		if l[i] == x {
			return i, true
		}
	}
	return -1, false
}

func main() {
	hit, psHits := checkMasterMind([]Slot{G, G, R, R})
	fmt.Printf("%d, %d \n", hit, psHits)

}
