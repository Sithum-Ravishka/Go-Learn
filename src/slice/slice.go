package main

import (
	"fmt"
	"unsafe"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix[0][2] = 20

	fmt.Println(matrix[1][0] * matrix[2][0])

	// a := [5]int{1, 2, 3, 4, 5} //40
	// b := []int{1, 2, 3, 4, 5}  //24

	k := [5]string{"a", "a", "a", "a", "a"} //80
	j := []string{"a", "a", "a", "a", "a"}  //24

	fmt.Printf("Array, Size bytes %d \n", unsafe.Sizeof(k))
	fmt.Printf("Slice, Size bytes %d \n", unsafe.Sizeof(j))
}
