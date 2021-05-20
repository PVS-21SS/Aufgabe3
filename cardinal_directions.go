package main

type Direction int

const (
	North Direction = 0
	East            = 1
	South           = 2
	West            = 3
)

func (d Direction) String() string {
	dir := []string{"North", "East", "South", "West"}
	if int(d) >= cap(dir) {
		d = Direction(int(d) - cap(dir))
	}
	return dir[d]
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

func directionCounter() int {
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
