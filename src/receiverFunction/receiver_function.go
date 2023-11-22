package main

import "fmt"

// Define a type named `Person`
type Person struct {
	FirstName string
	LastName  string
}

// Define a method with a pointer receiver
func (p *Person) UpdateLastName(newLastName string) {
	p.LastName = newLastName
}

func main() {
	// Create a new Person
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	// Print the initial state
	fmt.Println("Before Update:", person)

	// Call the method with a pointer receiver
	person.UpdateLastName("Smith")

	// Print the updated state
	fmt.Println("After Update:", person)
}
