package main

import "fmt"

type student struct {
	name  string
	age   int
	class int
	grade int
}

func main() {
	var std1 student

	std1.name = "Sithum"
	std1.age = 23
	std1.class = 9
	std1.grade = 99

	printStudent(std1)
	printStudent(std1)

}

func printStudent(std student) {
	fmt.Println("Student name: ", std.name)
	fmt.Println("Student age: ", std.age)
	fmt.Println("Student class: ", std.class)
	fmt.Println("Student grade: ", std.grade)
}
