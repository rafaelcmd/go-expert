package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 4
	myArray[2] = 10

	fmt.Println(len(myArray))

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for i, v := range myArray {
		fmt.Printf("The value of index %d is %d\n", i, v)
	}
}
