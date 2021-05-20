package main

type Axis int

const (
	NS Axis = iota
	EW
)

func (a Axis) String() string {
	axis := []string{"NorthSouth", "EastWest"}
	return axis[a]
}

func (a Axis) next() Axis {
	if a < 1 {
		return a+1
	}
	return 0
}
