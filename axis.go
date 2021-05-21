package main

type Axis int

func (a Axis) String() string {
	axis := []string{"NorthSouth", "EastWest"}
	return axis[a]
}

func (a Axis) next() Axis {
	if a < Axis((directionCounter()/2)-1) {
		return a+1
	}
	return 0
}
