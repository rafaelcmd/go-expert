package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Start processing request")
	defer log.Println("Finish processing request")
	select {
	case <-time.After(2 * time.Second):
		log.Println("Done processing request")
	case <-ctx.Done():

		log.Println("Request canceled by client")
	}
}
