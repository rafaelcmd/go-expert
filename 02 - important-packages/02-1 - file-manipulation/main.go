package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	f2, err := os.Create("test2.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.WriteString("This is a test file for GoLang file manipulation package")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", size)

	size, err = f2.Write([]byte("This is a test file for GoLang file manipulation package"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", size)

	f.Close()

	f3, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f3))

	// reading file with bufio to read line by line
	f4, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f4)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}

	err = os.Remove("test2.txt")
	if err != nil {
		panic(err)
	}
}
