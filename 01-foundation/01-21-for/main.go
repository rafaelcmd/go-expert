package main

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	names := []string{"John", "Paul", "George", "Ringo"}

	// for loop with range
	for k, v := range names {
		println(k, v)
	}

	// while loop
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// infinite loop
	for {
		println("loop")
	}
}
