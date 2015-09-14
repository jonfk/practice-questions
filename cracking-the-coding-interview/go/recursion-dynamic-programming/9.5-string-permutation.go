package main

import (
	"fmt"
)

func main() {
	test := "abcd"

	fmt.Printf("%v\n", permutate(test))
}

func permutate(a string) []string {
	if len(a) <= 1 {
		return []string{a}
	}
	asRune := []rune(a)
	letter := asRune[0]
	lessPermutated := permutate(string(asRune[1:]))
	var result []string
	for i := range lessPermutated {
		less := []rune(lessPermutated[i])
		for i := 0; i <= len(less); i++ {
			var temp []rune
			temp = append(temp, less[0:i]...)
			//fmt.Printf("less %v %d\n", string(temp), i)
			temp = append(temp, letter)
			temp = append(temp, less[i:]...)
			result = append(result, string(temp))
		}
	}
	return result
}
