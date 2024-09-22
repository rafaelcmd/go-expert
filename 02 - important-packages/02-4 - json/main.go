package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int     `json:"number"`
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{
		Number:  1001,
		Balance: 1000.50,
	}

	// Marshal returns the JSON encoding of account.
	result, error := json.Marshal(account)
	if error != nil {
		panic(error)
	}
	println(string(result))

	// NewEncoder returns a new encoder that writes to os.Stdout.
	error = json.NewEncoder(os.Stdout).Encode(account)
	if error != nil {
		panic(error)
	}

	rawJson := []byte(`{"number":1002,"balance":5000.50}`)
	var anotherAccount Account
	// Unmarshal parses the JSON-encoded data and stores the result in anotherAccount.
	error = json.Unmarshal(rawJson, &anotherAccount)
	if error != nil {
		panic(error)
	}
	println(anotherAccount.Number)
	fmt.Println(anotherAccount.Balance)
}
