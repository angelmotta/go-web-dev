package main

import (
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Firstname string
	Lastname string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	myUser := user{
		Firstname: "Angel",
		Lastname: "Motta",
	}

	index := func (w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
		HandleError(w, err)
	}

	dog := func (w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
		HandleError(w, err)
	}

	me := func (w http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(w, "me.gohtml", myUser)
		HandleError(w, err)
	}

	// Using Handle
	// http.HandlerFunc(f) allow the use of ordinary function 'f' as HTTP handlers.
	// if 'index' is a function with appropriate signature, HandlerFunc(index) is a Handler that calls index
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}