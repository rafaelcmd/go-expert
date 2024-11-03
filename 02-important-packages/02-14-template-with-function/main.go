package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	tmp := template.New("content.html")
	tmp.Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	})
	tmp = template.Must(tmp.ParseFiles(templates...))

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
