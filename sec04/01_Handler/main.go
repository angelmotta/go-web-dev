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

}
