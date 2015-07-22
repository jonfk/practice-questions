package main

import (
	"./common"
	"fmt"
	"sync"

	// removed time to get to deadlock faster
	//"time"
)

var wg, start *sync.WaitGroup

func main() {
	var forks []*common.Fork
	wg = new(sync.WaitGroup)
	start = new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		fork := common.NewFork(i)
		forks = append(forks, fork)
	}

	//fmt.Printf("%#v\n", forks)
	start.Add(1)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		fmt.Println("philosophers")
		if i == 4 {
			go philosopher(i, forks[i], forks[0])
		} else {
			go philosopher(i, forks[i], forks[i+1])
		}
	}
	start.Done()

	fmt.Println("Waiting")
	wg.Wait()
	fmt.Println("Done")
}

func philosopher(id int, left, right *common.Fork) {
	start.Wait()

	for {
		fmt.Printf("%d thinking\n", id)
		// think
		//time.Sleep(100 * time.Millisecond)
		// eat
		fmt.Printf("%d about to eat\n", id)
		// think
		left.Lock.Lock()
		right.Lock.Lock()
		fmt.Printf("%d eating\n", id)
		//time.Sleep(100 * time.Millisecond)

		// finish eating
		fmt.Printf("%d finished eating\n", id)
		left.Lock.Unlock()
		right.Lock.Unlock()
	}

	//wg.Done()
}
