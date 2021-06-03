package main

type CardinalDirection int

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

//returns a String of the cardinal Direction
func (d CardinalDirection) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
