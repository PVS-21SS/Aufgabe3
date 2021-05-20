package main

import "fmt"

type Direction int
type Colour int
type Axis int

const (
	North Direction = 0
	East            = 1
	South           = 2
	West            = 3
)

const (
	Red Colour = 0
	Yellow
	Green
)

const (
	NS Axis = 0
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
func (a Axis) String() string {
	axis := []string{"NorthSouth", "EastWest"}
	if int(a) >= cap(axis) {
		a = Axis(int(a) - cap(axis))
	}
	return axis[a]
}

func (d Direction) next() Direction {
	return d + 1
}

func (d Direction) opposite() Direction {
	return d + 2
}
func (d Direction) whichAxis() Axis {
	return Axis(int(d % 2))
}

func directions() int {
	var d = Direction(0)
	var dcpy = d
	var cnt = 1
	for {
		d = d.next()
		if d.String() == dcpy.String() {
			break
		}
		cnt++
	}
	return cnt
}

func (c Colour) next() Colour {
	return c + 1
}

func (a Axis) next() Axis {
	return a + 1
}

type TrafficLight struct {
	dir Direction
	col Colour
	ax Axis
}
func (t TrafficLight) String() string {
	var s = "["+t.dir.String() + ": "+ t.col.String()+ "]"
	return s
}

func main() {
	var d = Direction(0)
	var c = Colour(0)
	var cnt = directions()

	fmt.Println("First Direction: \t\t", d)
	fmt.Println("Next direction: \t\t", d.next())
	fmt.Println( "opposite direction: \t", d.opposite())
	fmt.Println("Axis: \t\t\t\t\t", d.whichAxis(), )

	fmt.Println()

	fmt.Println("First Colour:\t\t\t", c)
	fmt.Println("Next Colour:\t\t\t", c.next())


	fmt.Println("Count directions:\t\t", cnt)

	fmt.Println()

	t := []TrafficLight{}
	for i := 0; i < cnt; i++{
		t = append(t, TrafficLight{
			dir: d,
			col: c,
			ax: d.whichAxis(),})
		d = d.next()
	}
	fmt.Println(t)

}
