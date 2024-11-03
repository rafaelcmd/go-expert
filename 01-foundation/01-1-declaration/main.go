package main

// Package level declaration
const a = "Hello World"

// Multiple declaration
var (
	b bool    = true
	c int     = 10
	d string  = "Hello"
	e float64 = 1.2
)

func main() {
	//Short declaration operator
	a := "X"
	a = "Y"
	println(a)
}
