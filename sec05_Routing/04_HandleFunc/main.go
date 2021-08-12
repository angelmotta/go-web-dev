package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	fmt.Println("dog")
	io.WriteString(w, "new dog dog dog from HandleFunc")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "new cat cat cat from HandleFunc")
}

func main() {
	// Using HandleFunc
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	// Using Handle (as before)
	//http.Handle("/dog", http.HandleFunc(d))  // 2nd arg casting to type Handler
	//http.Handle("/dog", http.HandleFunc(c))
	// this is similar to this:
	// https://play.golang.org/p/X2dlgVSIrd
	// ---and this---
	// https://play.golang.org/p/YaUYR63b7L

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
