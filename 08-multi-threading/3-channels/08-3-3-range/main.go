package main

import "fmt"

func main() {
	ch := make(chan int)

	go publish(ch)
	receive(ch)
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // close the channel
}

func receive(ch chan int) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
	}
}
