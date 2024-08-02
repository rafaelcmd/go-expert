package main

import "fmt"

type Customer struct {
	name   string
	age    int
	active bool
	// Struct composition
	Address
}

type Address struct {
	street string
	city   string
	state  string
	zip    string
}

func main() {
	customer := Customer{
		name:   "John",
		age:    30,
		active: true,
		Address: Address{
			street: "123 Main St",
			city:   "Anytown",
			state:  "CA",
			zip:    "12345",
		},
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t, Street: %s, City: %s, State: %s, Zip: %s\n",
		customer.name, customer.age, customer.active, customer.street, customer.city, customer.state, customer.zip)
}
