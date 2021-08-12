package main

import (
	"log"
	"os"
	"text/template"
)


type hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		{
			"MyHotel1",
			"1551",
			"San Francisco",
			"CA1551",
			"Central",
		},{
			"MyHotel2",
			"3003",
			"Tenderloin",
			"CA94109",
			"South",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
