package main

import (
	"log"
	"os"
	"text/template"
)


type menu struct {
	Breakfast, Lunch, Dinner string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menus := []menu{
		{
			"Milk with bread",
			"Rust Potato",
			"Milk with cookies",
		},{
			"Coffee with fried eggs",
			"Rice and chicken",
			"Salad bar",
		},
	}

	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}
