package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// Implement Handler interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Angelinux-key", "this value is from Angelinux Header")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any content you want in this func</h1>")
}

func main() {
	var d hotdog // d is of type hotdog and Handler and the underlying type is int
	http.ListenAndServe(":8080", d)
}
