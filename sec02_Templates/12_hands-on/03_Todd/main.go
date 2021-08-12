package main

import (
	"log"
	"os"
	"text/template"
)


type item struct {
	Name, Description	string
	Price				float64
}

type meal struct {
	Meal string	// Breakfast, Lunch, Dinner
	Item []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	menus := menu{
		{
			Meal: "Breakfast",
			Item: []item{
				item{
					Name:    "Oatmeal",
					Description: "yum yum",
					Price:   4.95,
				},
				item{
					Name:    "Cheerios",
					Description: "American eating food traditional now",
					Price:   3.95,
				},
			},
		},{
			"Lunch",
			[]item{
				{
					Name:    "Hamburger",
					Description: "Delicious good eating for you",
					Price:   6.95,
				},
				{
					Name:    "Cheese Melted Sandwich",
					Description: "Make cheese bread melt grease hot",
					Price:   3.95,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}
