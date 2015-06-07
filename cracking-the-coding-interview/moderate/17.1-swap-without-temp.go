package main

import (
	"fmt"
)

func main() {
	a, b := 100, 223

	a = a + b

	b = a - b

	a = a - b

	fmt.Printf("%d %d \n", a, b)
}
