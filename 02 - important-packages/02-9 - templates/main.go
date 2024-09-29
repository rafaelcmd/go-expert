package main

import (
	"os"
	"text/template"
)

func main() {
	course := Course{"Golang", 10}
	tmp := template.New("template")
	tmp, _ = tmp.Parse("The course {{.Name}} will take {{.Time}} hours to complete.")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

type Course struct {
	Name string
	Time int
}
