package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("worker", id, "step", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	for id := 1; id <= 5; id++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println("worker", id)
		}(id)
	}

	wg.Wait()

	fmt.Println("all workers done")
}
