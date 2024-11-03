package main

import (
	"os"
	"text/template"
)

func main() {
	course := Course{"Golang", 10}
	tmp := template.Must(template.New("template").Parse("The course {{.Name}} will take {{.Time}} hours to complete."))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

type Course struct {
	Name string
	Time int
}
