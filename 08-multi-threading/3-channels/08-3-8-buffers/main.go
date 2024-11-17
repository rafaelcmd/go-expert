package main

func main() {
	ch := make(chan string, 2) // buffered channel

	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
