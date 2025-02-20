package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d has started their work\n", i)
	if i == 2 {
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Printf("Worker %d has ended their work\n", i)
}

func main() {
	fmt.Println("This is sync waitgroup demo:")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("This is sync waitgroup demo ended!!")
}
