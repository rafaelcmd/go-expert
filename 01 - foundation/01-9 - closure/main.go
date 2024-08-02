package main

import "fmt"

func main() {
	// Closure is a function that captures the variables from the context in which it was created
	total := func() int {
		return sum(1, 3, 5, 10, 21) * 2
	}()
	fmt.Println(total)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
