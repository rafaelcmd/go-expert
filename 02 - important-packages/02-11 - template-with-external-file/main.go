package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	tmp := template.Must(template.New("content.html").ParseFiles("content.html"))
	courses := Courses{
		Course{"Golang", 10},
		Course{"Python", 20},
		Course{"Java", 30},
	}
	err := tmp.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}
}
