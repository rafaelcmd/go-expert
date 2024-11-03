package main

import "fmt"

func main() {
	// Empty interfaces can hold values of any type
	var x interface{} = 10
	var y interface{} = "hello"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("The variable type is %T and the value is %v\n", t, t)
}
