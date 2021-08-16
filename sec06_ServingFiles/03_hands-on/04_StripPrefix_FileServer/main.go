package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", dog)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w,"index.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("template didn't execute: ", err)
	}
}