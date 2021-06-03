package main


type Axis struct {
	dirA    CardinalDirection
	dirB    CardinalDirection
	Channel chan CardinalDirection
}

//Returns the name of axis. The name is a combination of both cardinalDirections
func (a Axis) String() string {
	return a.dirA.String() + a.dirB.String()
}