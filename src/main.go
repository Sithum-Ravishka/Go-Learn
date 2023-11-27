package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Area() float64 {
	return float64(r.width * r.height)
}

func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {

	rect := Rectangle{width: 5, height: 10}

	printArea(rect)
}
