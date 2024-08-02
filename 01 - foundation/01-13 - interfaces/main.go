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

type Company struct {
	name   string
	active bool
	Address
}

// Interface is a set of method signatures
type Person interface {
	disable()
}

func (c Customer) disable() {
	c.active = false
	fmt.Printf("Customer has been disabled\n")
	fmt.Printf("The customer status is: %t\n", c.active)
}

func (c Company) disable() {
	c.active = false
	fmt.Printf("Company has been disabled\n")
	fmt.Printf("The company status is: %t\n", c.active)
}

func deactivate(person Person) {
	person.disable()
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

	deactivate(customer)

	company := Company{
		name:   "Acme Inc",
		active: true,
		Address: Address{
			street: "456 Oak St",
			city:   "Anytown",
			state:  "CA",
			zip:    "12345",
		},
	}

	deactivate(company)
}
