package main

import (
	"fmt"
	"sync"
)

func main() {
	scores := make(map[int]int)

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			scores[id] = id * 10
		}(i)
	}

	wg.Wait()

	fmt.Println("map size:", len(scores))
	fmt.Println("score 1:", scores[1])
	fmt.Println("score 100:", scores[100])
}
