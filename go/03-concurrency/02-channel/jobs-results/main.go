package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("worker", id, "received job", job)

		time.Sleep(300 * time.Millisecond)

		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	for workerID := 1; workerID <= 3; workerID++ {
		wg.Add(1)
		go worker(workerID, jobs, results, &wg)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("result:", result)
	}

	fmt.Println("all results received")
}
