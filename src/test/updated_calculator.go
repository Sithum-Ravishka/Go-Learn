package main

import (
	"fmt"
	"os"
)

type Calculator interface {
	UserInput() error
	Calculation(firstNo int, operator string, secondNo int) (int, error)
}

type SimpleCalculator struct {
	Total int
}

const (
	OperatorAdd      = "+"
	OperatorSubtract = "-"
	OperatorMultiply = "*"
	OperatorDivide   = "/"
	OperatorEqual    = "="
)

func main() {
	calculator := SimpleCalculator{}
	err := calculator.UserInput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (c *SimpleCalculator) UserInput() error {
	var (
		firstNo  int
		secondNo int
		operator string
	)

	fmt.Println("Enter First Number: ")
	_, err := fmt.Scan(&firstNo)
	if err != nil {
		return err
	}

	c.Total = firstNo

	for {
		fmt.Println("Enter Calculation Operator(+, -, *, /) or =: ")
		_, err := fmt.Scan(&operator)
		if err != nil {
			return err
		}

		if operator == OperatorEqual {
			fmt.Println("Total Result is:", c.Total)
			break
		}

		if isValidOperator(operator) {
			fmt.Println("Enter Second Number: ")
			_, err := fmt.Scan(&secondNo)
			if err != nil {
				return err
			}

			result, err := c.Calculation(c.Total, operator, secondNo)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Current Result:", result)
			c.Total = result
		} else {
			fmt.Println("Invalid Operator. Please enter a valid operator(+, -, *, / or exit for =)")
		}
	}

	return nil
}

func (c *SimpleCalculator) Calculation(firstNo int, operator string, secondNo int) (int, error) {
	var result int

	switch operator {
	case OperatorAdd:
		result = firstNo + secondNo
	case OperatorSubtract:
		result = firstNo - secondNo
	case OperatorMultiply:
		result = firstNo * secondNo
	case OperatorDivide:
		if secondNo != 0 {
			result = firstNo / secondNo
		} else {
			return 0, fmt.Errorf("can't divide by zero")
		}
	default:
	}

	return result, nil
}

func isValidOperator(operator string) bool {
	return operator == OperatorAdd || operator == OperatorSubtract || operator == OperatorMultiply || operator == OperatorDivide
}
