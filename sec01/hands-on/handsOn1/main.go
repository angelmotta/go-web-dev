package main

import (
	"fmt"
	"math"
)

type Square struct {
	edge float64
}

type Circle struct {
	radius float64
}

func (s Square) getArea() float64 {
	fmt.Println(" ---- getArea() method Square ----")
	return s.edge * s.edge
}

func (c Circle) getArea() float64 {
	fmt.Println(" ---- getArea() method Circle ----")
	return math.Pi * math.Pow(c.radius, 2)
}

type Shape interface {
	getArea() float64
}

// Receive 'interface' figure
func info(figure Shape) {
	fmt.Println(" --- Info ---")
	fmt.Println(figure.getArea())
}

func main () {
	mySquare := Square {4}
	myCircle := Circle {3}

	info(mySquare)	// mySquare is also an interface
	info(myCircle)	// myCircle is also an interface
}
