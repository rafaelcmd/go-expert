package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("content.html").ParseFiles("content.html"))
		courses := Courses{
			Course{"Golang", 10},
			Course{"Python", 20},
			Course{"Java", 30},
		}
		err := tmp.Execute(w, courses)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8082", nil)
}
