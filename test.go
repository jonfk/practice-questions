package main

import (
	"bytes"
	"fmt"
)

// []  “0-99”
// [1,99] “0,2-98”
// [2,3,4,5,98,99] “0-1,6-97”
// [99] “0-98”
// [98] “0-97,99”
// [0] “1-99”
// [1] “0,2-99”

func missingNums(nums []int) string {

	if len(nums) < 1 {
		return "0-99"
	}
	buf := new(bytes.Buffer)

	var prev int = -1
	for i := range nums {
		if prev == (nums[i] - 1) {
			prev = nums[i]
		} else if prev+1 == nums[i]-1 {
			buf.WriteString(fmt.Sprintf("%d", prev+1))
			prev = nums[i]
			buf.WriteString(",")
		} else {
			buf.WriteString(fmt.Sprintf("%d-%d", prev+1, nums[i]-1))
			prev = nums[i]
			buf.WriteString(",")
		}
		if i == len(nums)-1 && nums[i] != 99 {
			if nums[i]+1 == 99 {
				buf.WriteString(fmt.Sprintf("%d", 99))
			} else {
				buf.WriteString(fmt.Sprintf("%d-%d", nums[i]+1, 99))
			}
			break
		} else if i == len(nums)-1 {
			break
		}
	}
	// last comma
	result := buf.String()
	if string([]rune(result)[len(result)-1]) == "," {
		return result[:len(result)-1]
	}
	return buf.String()
}

func main() {
	fmt.Printf("%s\n", missingNums([]int{}))
	fmt.Printf("%s\n", missingNums([]int{1, 99}))
	fmt.Printf("%s\n", missingNums([]int{2, 3, 4, 5, 98, 99}))
	fmt.Printf("%s\n", missingNums([]int{99}))
	fmt.Printf("%s\n", missingNums([]int{98}))
	fmt.Printf("%s\n", missingNums([]int{0}))
	fmt.Printf("%s\n", missingNums([]int{1}))
}
