package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", blog{"My Blog"})

	go func() {
		http.ListenAndServe(":8080", mux)
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Mux2!"))
	})

	go func() {
		http.ListenAndServe(":8081", mux2)
	}()

	select {}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
