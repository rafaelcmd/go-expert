package main

import "fmt"

func main() {
	msg := make(chan string)
	go receive("Hello", msg)
	read(msg)
}

func receive(content string, message chan<- string) { // send-only channel
	message <- content
}

func read(data <-chan string) { // receive-only channel
	fmt.Println(<-data)
}
