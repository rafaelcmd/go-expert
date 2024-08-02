package main

import "fmt"

// Struct is a collection of fields
type Customer struct {
	name   string
	age    int
	active bool
}

func main() {
	customer := Customer{
		name:   "John",
		age:    30,
		active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", customer.name, customer.age, customer.active)
}
