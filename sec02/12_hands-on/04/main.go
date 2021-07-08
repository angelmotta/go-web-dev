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

type restaurant struct {
	Name	string
	Menu 	[]meal
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	restaurantsList := restaurants{
		restaurant{
			Name: "The good restaurant",
			Menu: []meal{
				{
					Meal: "Breakfast",
					Item: []item{
						{
							Name:        "soup",
							Description: "descr1",
							Price:       4.15,
						},
						{
							Name:        "bread and milk",
							Description: "descr2",
							Price:       5.12,
						},
					},
				},
				{
					Meal: "Lunch",
					Item: []item{
						{
							Name:        "Hamburger",
							Description: "Delicious good eating for you",
							Price:       6.95,
						},
						{
							Name:        "Cheese Melted Sandwich",
							Description: "Make cheese bread melt grease hot",
							Price:       3.95,
						},
					},
				},
			},
		},
	}

	myNewRestaurant := restaurant{
		Name: "The New Restaurant",
		Menu: []meal {
			{
				Meal: "Breakfast",
				Item: []item{
					{
						Name:    "Oatmeal",
						Description: "yum yum",
						Price:   4.95,
					},
				},
			},
			{
				Meal: "Lunch",
				Item: []item{
					{
						Name: "Vegatables",
						Description: "healthy",
						Price: 6.99,
					},
				},
			},
		},
	}


	// My Slice
	resSlice := make([]restaurant, 0)
	for _, res := range restaurantsList {
		resSlice = append(resSlice, res)
	}

	// Add new restaurant
	resSlice = append(resSlice, myNewRestaurant)

	// Execute template
	err := tpl.Execute(os.Stdout, resSlice)
	if err != nil {
		log.Fatalln(err)
	}
}
