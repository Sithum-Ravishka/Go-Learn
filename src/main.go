package main

import "fmt"

// Define a type named `Person`
type Person struct {
	FirstName string
	LastName  string
}

// Define a method named `FullName` with a receiver of type `Person`
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	// Create a new Person
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	// Call the method using the dot notation
	fullName := person.FullName()

	// Print the result
	fmt.Println("Full Name:", fullName)
}
