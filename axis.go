package main

type Axis int

const (
	NS Axis = iota
	EW
)

func (a Axis) String() string {
	axis := []string{"NorthSouth", "EastWest"}
	if int(a) >= cap(axis) {
		a = Axis(int(a) - cap(axis))
	}
	return axis[a]
}

func (a Axis) next() Axis {
	return a + 1
}
