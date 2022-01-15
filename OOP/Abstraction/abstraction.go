package main

import "fmt"

// Define a new data type "Triangle"
type Triangle struct {
	base, height float32
}

// Define a new data type "Square"
type Square struct {
	length float32
}

// Define a new data type "Rectangle"
type Rectangle struct {
	length, width float32
}

// Define a new data type "Circle"
type Circle struct {
	radius float32
}

// A method for type "Triangle"
func (t Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}

// A method for type "Square"
func (l Square) Area() float32 {
	return l.length * l.length
}

// A method for type "Rectangle"
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

// A method for type "Circle"
func (c Circle) Area() float32 {
	return 3.14 * (c.radius * c.radius)
}

// Define an interface as achieve abstraction
type Shape interface {
	Area() float32
}

func main() {
	// Declare and assign values to varaibles
	t := Triangle{base: 15, height: 25}
	s := Square{length: 5}
	r := Rectangle{length: 5, width: 10}
	c := Circle{radius: 5}

	// Define a variable of type interface
	var sh Shape

	// Assign to the interface a variable of type "Triangle"
	sh = t
	fmt.Println("Area of Triangle", sh.Area())

	// Assign to the interface a variable of type "Square"
	sh = s
	fmt.Println("Area of Square", sh.Area())

	// Assign to the interface a variable of type "Rectangle"
	sh = r
	fmt.Println("Area of Rectangle", sh.Area())

	// Assign to the interface a variable of type "Circle"
	sh = c
	fmt.Println("Area of Circle", sh.Area())
}
