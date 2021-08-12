package main

import "fmt"

// Composition in action (inheritance behaviour)

type Person struct {
	FirstName string
	LastName string
	Age int
}

type SecretAgent struct {
	Person		// Composition in action
	licenseToKill bool
}

// Methods definition
func (p Person) pSpeak () {
	fmt.Println("Hi, I am person. My name is", p.FirstName)
}

func (sa SecretAgent) saSpeak () {
	fmt.Println("Hi, I am Secret Agent. My name is", sa.FirstName)
	if sa.licenseToKill {
		fmt.Println("I am authorized to kill")
	} else {
		fmt.Println("I am Not authorized to kill")
	}
}
// Implementing Stringer interface (implicitly)
func (p Person) String() string {
	return fmt.Sprintf("'Stringer' interface implemented | My name: %s", p.FirstName)
}

func main()  {
	myPerson1 := Person{"Angel", "Motta", 36}
	mySecretAgent1 := SecretAgent{Person{"James", "Bond", 45}, true}
	fmt.Println(myPerson1.FirstName)
	myPerson1.pSpeak()

	fmt.Println(mySecretAgent1.FirstName)
	mySecretAgent1.saSpeak()
	mySecretAgent1.pSpeak()	// Composition works like inheritance

	// Print objects
	fmt.Println(myPerson1)		// Interface String implemented
	fmt.Println(mySecretAgent1)	// secretAgent also implemented Interface (inheritance using Composition)
}
