package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.Handle("/blog", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Blog"))
	}))
	http.ListenAndServe(":8080", mux)
}
