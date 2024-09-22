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
	result, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))
	err = req.Body.Close()
	if err != nil {
		return
	}
}
