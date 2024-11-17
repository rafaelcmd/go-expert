package main

// thread 1
func main() {
	ch := make(chan string) // empty channel

	// thread 2
	go func() {
		ch <- "Hello" // full channel
	}()

	// thread 1
	msg := <-ch // empty channel
	println(msg)
}
