package main

import "fmt"

func main() {
	UserInput()
}

func UserInput() {
	var (
		firstNo  int
		secondNo int
		operator string
	)

	fmt.Println("Enter First Number: ")
	fmt.Scan(&firstNo)
	total := firstNo

	for {

		fmt.Println("Enter Calculation Operator(+,-,*,/) or =: ")
		fmt.Scan(&operator)

		if operator == "=" {
			fmt.Println("Total Result is:  ", total)
			break
		}

		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			fmt.Println("Enter Second Number: ")
			fmt.Scan(&secondNo)

			// Return value equal to total for continue calculation
			total = Calculation(total, operator, secondNo)
			fmt.Println("Current Result:", total)
		} else {
			fmt.Println("Invalid Operator. Please enter valid operator(+,-,*,/ or exit for =")
		}
	}
}

// Calculation Part
func Calculation(firstNo int, operator string, secondNo int) int {
	var result int

	if operator == "+" {
		result = firstNo + secondNo
	} else if operator == "-" {
		result = firstNo - secondNo
	} else if operator == "*" {
		result = firstNo * secondNo
	} else if operator == "/" {
		if secondNo != 0 {
			result = firstNo / secondNo
		} else {
			fmt.Println("Can't division by zero.")
		}
	} else {
		fmt.Println("Invalid Operator.")
	}

	return result
}
