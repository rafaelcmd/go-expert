package main

import "fmt"

func main() {
	var myVar interface{} = "Hello, World!"
	println(myVar.(string))

	res, ok := myVar.(int)
	fmt.Printf("The value of res is %v and the result of ok is %v\n", res, ok)
}
