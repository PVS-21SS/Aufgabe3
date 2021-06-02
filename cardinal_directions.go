package main

type CardinalDirection int
const (
	north = iota
	east
	south
	west
)
func (direction CardinalDirection) String() string {
	return [...]string{"North", "East", "South", "West"}[direction]
}