package main

import (
	"fmt"
	sc "strconv"
)

// Base struct
type Vehicle struct {
	tyres int
	seats int
	name  string
}

// Base struct method inherited by all compositions
func (v *Vehicle) Usage() {
	fmt.Println("This vehicle has " + sc.Itoa(v.tyres) + " tyres, can accommodate " + sc.Itoa(v.seats) + " people and called as " + v.name)
}

// Car extends Vehicle
type Car struct {
	// Composition - Embedding another struct inside struct.
	Vehicle
}

// Motercycle extends Vehicle
type Motercycle struct {
	// Composition - Embedding another struct inside struct.
	Vehicle
}

// Auto extends Vehicle
type Auto struct {
	// Composition - Embedding another struct inside struct.
	Vehicle
}

func main() {
	c := Car{
		Vehicle{
			tyres: 4,
			seats: 5,
			name:  "Car",
		},
	}
	c.Usage()
	a := Auto{
		Vehicle{
			tyres: 3,
			seats: 3,
			name:  "Auto",
		},
	}
	a.Usage()
	m := Motercycle{
		Vehicle{
			tyres: 2,
			seats: 2,
			name:  "Motercylce",
		},
	}
	m.Usage()
}
