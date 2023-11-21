package main

import "fmt"

type Student struct {
	name  string
	age   int
	class int
	grade int
}

func main() {
	var std1 Student

	fmt.Println("Enter student name:")
	fmt.Scanln(&std1.name)

	fmt.Println("Enter student age:")
	fmt.Scanln(&std1.age)

	fmt.Println("Enter student class:")
	fmt.Scanln(&std1.class)

	fmt.Println("Enter student grade:")
	fmt.Scanln(&std1.grade)

	printStudent(std1)

}

func printStudent(std Student) {
	fmt.Println("Student name: ", std.name)
	fmt.Println("Student age: ", std.age)
	fmt.Println("Student class: ", std.class)
	fmt.Println("Student grade: ", std.grade)
}
