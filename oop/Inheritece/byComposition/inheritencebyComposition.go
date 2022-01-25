package inheritence

import (
	"fmt"
	sc "strconv"
)

// Base struct
type Vehicle struct {
	tires int
	seats int
	name  string
}

// Base struct method inherited by all compositions
func (v *Vehicle) Usage() {
	fmt.Println("This vehicle has " + sc.Itoa(v.tires) + " tires, can accommodate " + sc.Itoa(v.seats) + " people and called as " + v.name)
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
			tires: 4,
			seats: 5,
			name:  "Car",
		},
	}
	c.Usage()
	a := Auto{
		Vehicle{
			tires: 3,
			seats: 3,
			name:  "Auto",
		},
	}
	a.Usage()
	m := Motercycle{
		Vehicle{
			tires: 2,
			seats: 2,
			name:  "Motercylce",
		},
	}
	m.Usage()
}
