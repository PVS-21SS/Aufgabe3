package main

type Axis int

// toString to Print the Axis and to define which Axis exist
func (a Axis) String() string {
	axis := []string{"NorthSouth", "EastWest"}
	return axis[a]
}

// Next Axis Function, based on the Index
func (a Axis) next() Axis {
	if a < Axis((directionCounter()/2)-1) {
		return a+1
	}
	return 0
}
