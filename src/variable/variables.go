//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor string = "Red"
	fmt.Println("Favorite Color is: ", favoriteColor)

	var year, age = 1998, 25
	fmt.Printf("Year is %d and age is %d \n", year, age)

	var (
		firstInitials = "R"
		lastInitials  = "M"
	)
	fmt.Println("Initials is: ", firstInitials, lastInitials)

	var days int
	days = age * 365

	fmt.Println("Your age (in days): ", days)
}
