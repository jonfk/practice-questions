package main

import (
	"fmt"
)

func fibN(n int) chan int {
	c := make(chan int)

	go func() {
		x, y := 0, 1

		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}

		close(c)
	}()

	return c
}

//It’s better than sending a variable to 'quit’ channel like,

//  quit <- 1 // stop the goroutine
//because you can stop more than one goroutines by closing it.

func fib() (chan int, chan int) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		x, y := 0, 1

		keepGoing := true
		for keepGoing {
			select {
			case <-quit:
				keepGoing = false
			default:
				c <- x
				x, y = y, x+y
			}
		}

		close(c)
	}()

	return c, quit

}

func main() {
	// for i := range fibN(10) {
	// 	fmt.Printf("%d ", i)
	// }
	c, quit := fib()
	for i := range c {
		fmt.Printf("%d ", i)
		if i == 55 {
			close(quit)
		}
	}
	fmt.Printf("\n")
}
