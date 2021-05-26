package main

type Colour int


// toString to Print the TrafficLight with the Colour Name
func (c Colour) String() string {
	col := []string{"Red", "Green", "Yellow"}
	return col[c]
}

// Next Colour Function, based on the Index
func (c Colour) next() Colour {
	if c < 2 {
		return c+1
	}
	return 0
}