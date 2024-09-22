package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	// Defer is a statement that schedules a function call to be run after the function completes.
	defer req.Body.Close()
	result, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))
}
