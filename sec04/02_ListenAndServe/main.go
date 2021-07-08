package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// hotdog implicitly implement the Handler Interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w,"My power code here!")
}

func main() {
	var myHotDogHandler hotdog

	// ListenAndServer receive receive a struct of type Handler (myHotDogHandler)
	http.ListenAndServe(":8080", myHotDogHandler)
}
