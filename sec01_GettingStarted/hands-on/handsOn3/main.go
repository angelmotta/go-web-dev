package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

type SecretAgent struct {
	Person		// Composition in action
	licenseToKill bool
}

func (p Person) haveFeelings() string{
	return fmt.Sprintf("I have feelings. My name is %s", p.FirstName)
}

// If SecretAgent does not implement this interface directly, inherit method from Person
func (sa SecretAgent) haveFeelings() string{
	return fmt.Sprintf("I have feelings. My license to kill not means nothing. My name is %s", sa.FirstName)
}

type Human interface {
	haveFeelings() string
}

// Receive an 'interface' as an argument
func tellMeSomething(p Human) {
	fmt.Println("Ok, I tell you something")
	switch myValue := p.(type) {
	case Person:
		fmt.Println("I am Person:", myValue.FirstName)
	case SecretAgent:
		fmt.Println("I am Secret Agent:", myValue.FirstName)
	default:
		fmt.Println("Unknown")
	}
	fmt.Println(p.haveFeelings())
}

func main()  {
	myPerson1 := Person{"Angel", "Motta", 36}
	mySecretAgent1 := SecretAgent{Person{"James", "Bond", 45}, true}
	// Print raw values from Objects (Stringer interface not implemented)
	fmt.Println(myPerson1)
	fmt.Println(mySecretAgent1)

	// Receive interface 'human' (Polymorphism in action)
	tellMeSomething(myPerson1)			// myPerson is also a human
	tellMeSomething(mySecretAgent1)		// mySecretAgent is also a human

	// Methods
	fmt.Println("-------- methods ---------")
	fmt.Println(mySecretAgent1.haveFeelings())
	fmt.Println(mySecretAgent1.Person.haveFeelings())
}
