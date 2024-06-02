package main

import (
	"errors"
	"fmt"
)

func main() {
	value, boolValue, err := sum(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value, boolValue, err)
}

func sum(a, b int) (int, bool, error) {
	if a+b >= 5 {
		return a + b, true, nil
	}
	return a + b, false, errors.New("Sum is less than 5.")
}
