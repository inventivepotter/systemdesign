package main

import "fmt"

// Defining new data type Counter
// Counter has an un-exported variable 'value'
type Counter struct {
	value int
}

// Counter has a method add() to manipulate it's variable.
func (c *Counter) add() {
	c.value = c.value + 1
}

// Counter has a method get() to get it's variable.
func (c *Counter) get() int {
	return c.value
}

// Given
func main() {
	c := Counter{
		value: 5,
	}
	c.add()
	fmt.Println(c.get())
}
