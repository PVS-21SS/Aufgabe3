package main

type CardinalDirection int
const (
	north = iota
	east
	south
	west
)

func (d CardinalDirection) String() string {
	dir := []string{"North", "East", "South", "West"}
	if int(d) >= cap(dir) {
		d = CardinalDirection(int(d) - cap(dir))
	}
	return dir[d]
}
