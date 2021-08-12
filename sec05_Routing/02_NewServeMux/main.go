package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

type myhome int

func (c myhome) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome home page")
}

func main() {
	var d hotdog
	var c hotcat
	//var zdef myhome

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	//mux.Handle("/", zdef)

	http.ListenAndServe(":8080", mux)
}
