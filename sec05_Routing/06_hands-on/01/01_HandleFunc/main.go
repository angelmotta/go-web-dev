package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	index := func (w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Index Page")
	}

	dog := func (w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Dog Page")
	}

	me := func (w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Angel Motta")
	}

	// Using HandleFunc
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
