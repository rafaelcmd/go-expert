package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s - Iteration %d\n", name, i)
		time.Sleep(1 * time.Second)
		waitGroup.Done()
	}
}

// thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	// thread 2
	go task("A", &waitGroup)
	// thread 3
	go task("B", &waitGroup)

	// thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Task C - Anonymous Iteration %d\n", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
