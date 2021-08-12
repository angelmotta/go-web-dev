package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

// hotdog implicitly implement the Handler Interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // enable access the 'Form' field from the body and also url parameters
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method			string
		URL				*url.URL
		Submissions		url.Values // map[string][]string
		Header			http.Header // map[string][]string
		Host			string
		ContentLength	int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var myHanlder hotdog
	http.ListenAndServe(":8080", myHanlder)
}
