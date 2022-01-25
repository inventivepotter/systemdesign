package main

import "fmt"

// Define an interface to achieve abstraction
type Polygon interface {
	Area() float32
}

// Define a new data type "Triangle"
type Triangle struct {
	base, height float32
}

// Define a new data type "Square"
type Square struct {
	length float32
}

// A method for type "Triangle" with parameter being t Trinagle
func (t Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}

// Same method for type "Square" with parametr being s square
func (s Square) Area() float32 {
	return s.length * s.length
}

func main() {
	// Declare and assign values to varaibles
	t := Triangle{base: 15, height: 25}
	s := Square{length: 5}

	// Define a variable of type interface
	var p Polygon

	// Assign to the interface a variable of type "Triangle"
	p = t
	// Ablet to call Area() method with different signitures
	fmt.Println("Area of Triangle", p.Area())

	// Assign to the interface a variable of type "Square"
	p = s
	fmt.Println("Area of Square", p.Area())
}
