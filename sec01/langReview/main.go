package main

import "fmt"

type hotint int

type person struct {
	fname string
	lname string
}

// Composition
type secretAgent struct {
	person
	licenseToKill bool
}

// Interface define functionality and allow polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("--- Speak 'method' ---")
	fmt.Println(p.fname, `says: "Good morning, Angel"`)
}

func (sa secretAgent) speak() {
	fmt.Println("--- Speak 'method' ---")
	if sa.licenseToKill {
		fmt.Println(sa.fname, sa.person.lname, `says: "Shaken, not stirred"`)
	} else {
		fmt.Println("No license to kill")
	}
}

func main () {
	// Composite literal
	xSlice := []int{2, 4, 6, 8}
	fmt.Printf("%T\n", xSlice)
	fmt.Println(xSlice)

	// Composite literal
	myMap := map[string]int{
		"Angel": 36,
		"Piero": 21,
	}
	fmt.Printf("%T\n", myMap)
	fmt.Println(myMap)

	var x hotint
	x = 10
	fmt.Println(x)

	p1 := person{
		"Rosa",
		"Paz",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()
	sa1.person.speak()

	// Polymorphism
	fmt.Println("\nPolymorphism in action")
	saySomething(p1)
	saySomething(sa1)
}
