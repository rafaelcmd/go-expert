package main

import "fmt"

type Customer struct {
	name   string
	age    int
	active bool
	Address
}

type Address struct {
	street string
	city   string
	state  string
	zip    string
}

// Method is a function with a special receiver argument
func (c Customer) disable() {
	c.active = false
	fmt.Printf("Customer has been disabled\n")
	fmt.Printf("The customer status is: %t\n", c.active)
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

	customer.disable()
}
