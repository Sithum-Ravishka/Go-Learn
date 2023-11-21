package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix[0][2] = 20

	fmt.Println(matrix[1][0] * matrix[2][0])
}
