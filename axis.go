package main

type Axis struct {
	dirA    CardinalDirection
	dirB    CardinalDirection
	Channel chan CardinalDirection
}

func (a Axis) String() string {
	return a.dirA.String() + a.dirB.String()
}