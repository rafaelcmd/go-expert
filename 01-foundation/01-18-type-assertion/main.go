package main

import "fmt"

func main() {
	var myVar interface{} = "Hello, World!"
	println(myVar.(string))

	// Type assertion with ok is used to check if the type assertion is successful
	res, ok := myVar.(int)
	fmt.Printf("The value of res is %v and the result of ok is %v\n", res, ok)
}
