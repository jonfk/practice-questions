package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	channel := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(done chan int) {
			wg.Add(1)
			var counter int
			for i := 0; i < 5000000; i++ {
				counter += 1
			}
			done <- counter
			wg.Done()
		}(channel)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		sum := <-channel
		fmt.Printf("%d\n", sum)
	}
	fmt.Println("done")
}
