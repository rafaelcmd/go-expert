package main

import "fmt"

func main() {
	// A slice is a dinamically-sized, flexible view into the elements of an array.
	// It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:3]), cap(s[:3]), s[:3])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 6)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}
