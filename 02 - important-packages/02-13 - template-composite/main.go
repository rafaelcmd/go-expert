package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	tmp := template.Must(template.New("content.html").ParseFiles(templates...))
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
