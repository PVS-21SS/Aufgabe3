package main

import "fmt"

type Direction int
type Colour int
type Axis int

const (
	North Direction = iota
	East
	South
	West
)

const (
	Red Colour = iota
	Yellow
	Green
)

const (
	NS Axis = iota
	EW
)

func (d Direction) String() string {
	dir := []string{"North", "East", "South", "West"}
	if int(d) >= cap(dir) {
		d = Direction(int(d) - cap(dir))
	}
	return dir[d]
}

func (c Colour) String() string {
	col := []string{"Red", "Yellow", "Green"}
	if int(c) >= cap(col) {
		c = Colour(int(c) - cap(col))
	}
	return col[c]
}

func dnext(d Direction) (Direction) {
	return d + 1
}

func dopposite(d Direction) Direction {
	return d + 2
}
func dmax(d Direction) Direction {
	return d
}

func cnext(c Colour) Colour {
	return c + 1
}
func anext(a Axis) Axis {
	return a + 1
}

func main() {
	var d = North
	var c = Green

	fmt.Println("Start:\t\t", c)
	fmt.Println("Next:\t\t", cnext(c))
	fmt.Println("---------------")
	fmt.Println("Start:\t\t", d)

	fmt.Println("Next:\t\t", dnext(d))
	fmt.Println("Opposite:\t", dopposite(d))
}
