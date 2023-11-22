package main

import "fmt"

// Define an interface named "Shape"
type Shape interface {
	// Method to calculate the area of the shape
	Area() float64
}

// Define a struct for a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define a struct for a circle
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Test struct {
	X float64
	Y float64
	Z float64
}

func (a Test) Area() float64 {
	return a.X + a.Y*a.Z
}

// Function that takes a Shape interface as a parameter
func PrintArea(shape Shape) {
	fmt.Printf("Area of the shape: %f\n", shape.Area())
}

func Interfac() { //use main
	// Create instances of Rectangle and Circle
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}
	test := Test{X: 20, Y: 10, Z: 30}

	// Call the PrintArea function with different shapes
	PrintArea(rectangle)
	PrintArea(circle)
	PrintArea(test)
}
