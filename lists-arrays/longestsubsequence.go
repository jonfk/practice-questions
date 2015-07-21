package main

import (
	"fmt"
)

func main() {
	tet := "abbacbccbcabccccccc"
	fmt.Printf("%s\n", subStrings(tet))
}

func subStrings(a string) string {
	if len(a) < 1 {
		return ""
	} else if len(a) == 1 {
		return a
	}

	var (
		maxSubString string
		maxLen       int
	)
	asRune := []rune(a)
	for i := 0; i < len(asRune); i++ {
		var charsSoFar map[rune]int = make(map[rune]int)
		for j := 1; j < len(asRune)-i; j++ {
			newChar := asRune[i+j-1]
			charsSoFar[newChar] += 1
			fmt.Printf("%s %d\n", string(asRune[i:i+j]), count(charsSoFar))
			if count(charsSoFar) <= 2 {
				continue
			} else {
				sub := string(asRune[i : i+j-1])
				if len(sub) > maxLen {
					maxSubString = sub
					maxLen = len(sub)
				}
				break
			}

		}
	}

	return maxSubString
}

func count(chars map[rune]int) int {
	count := 0
	for _, _ = range chars {
		count++
	}
	return count
}
