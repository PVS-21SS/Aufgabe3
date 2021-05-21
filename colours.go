package main

type Colour int

const (
	Red Colour = iota
	Yellow
	Green
)

// toString to Print the TrafficLight with the Colour Name
func (c Colour) String() string {
	col := []string{"Red", "Yellow", "Green"}
	return col[c]
}

// Next Colour Function, based on the Index
func (c Colour) next() Colour {
	if c < 2 {
		return c+1
	}
	return 0
}