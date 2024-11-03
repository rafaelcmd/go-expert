package main

func main() {
	a := 10
	b := 20

	// if statement
	if a < b {
		println("a is less than b")
	} else {
		println("a is not less than b")
	}

	//switch statement
	name := "John"

	switch name {
	case "John":
		println("John")
	case "Paul":
		println("Paul")
	case "George":
		println("George")
	case "Ringo":
		println("Ringo")
	default:
		println("default")
	}
}
