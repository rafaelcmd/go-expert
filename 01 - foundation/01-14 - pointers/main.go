package main

func main() {
	a := 10
	println(a)
	println(&a)

	// Pointer is a type that stores the memory address of another variable
	var pointer *int = &a
	println(pointer)

	*pointer = 20
	println(a)

	b := &a
	println(b)
	println(*b)

	*b = 30
	println(a)
}
