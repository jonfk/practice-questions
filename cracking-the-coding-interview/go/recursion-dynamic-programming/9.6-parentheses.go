package main

import "fmt"

func main() {
	fmt.Printf("%v\n", combinations(4))
}

func combinations(n int) []string {
	if n < 1 {
		return nil
	} else if n == 1 {
		return []string{"()"}
	} else {
		lessCombi := combinations(n - 1)
		var results []string
		for i := range lessCombi {
			less := []rune(lessCombi[i])
			paren := []rune("()")
			for j := 0; j <= len(less); j++ {
				var temp []rune
				temp = append(temp, less[0:j]...)
				temp = append(temp, paren...)
				temp = append(temp, less[j:]...)
				tempStr := string(temp)
				if !exists(results, tempStr) {
					results = append(results, tempStr)
				}
			}
		}
		return results
	}
}

func exists(arr []string, str string) bool {
	for i := range arr {
		if arr[i] == str {
			return true
		}
	}
	return false
}
