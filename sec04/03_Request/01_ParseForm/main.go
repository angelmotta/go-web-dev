package main

import (
	"log"
	"net/http"
	"text/template"
)

type hotdog int

// hotdog implicitly implement the Handler Interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // execute ParseForm to access the 'Form' field from the body
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form) // Form is a aap[string][]string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var myHanlder hotdog
	http.ListenAndServe(":8080", myHanlder)
}
