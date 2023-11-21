package main

import (
	"fmt"
)

func Variable() {
	var a, b, c, d int = 1, 2, 3, 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// block variable

func BlockVariable() {
	var (
		a int
		b int    = 1 // a := 1 can't be used in block variable
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// package main

// import (
// 	"fmt"
// )

func World() {
	var i, j string = "Hello", "World"

	fmt.Println(i, j) 
}

// using print with int value 10 20 space are auto add

func PrintInt() {
	var i, j = 10, 20

	fmt.Print(i, j)
}
