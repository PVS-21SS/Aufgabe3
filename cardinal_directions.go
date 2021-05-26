package main

type Direction int

// toString to Print the Direction and to define which directions exist
func (d Direction) String() string {
	dir := []string{"North", "East", "South", "West"}
	if int(d) >= cap(dir) {
		d = Direction(int(d) - cap(dir))
	}
	return dir[d]
}

// Next direction Function, based on the Index
func (d Direction) next() Direction {
	if d < Direction(directionCounter()-1) {
		return d + 1
	}
	return 0
}

// which axis is my input on
func (d Direction) whichAxis() Axis {
	return Axis(d % 2)
}

// counts the directions by iterating the String Array
func directionCounter() int {
	var d = Direction(0)
	var dCpy = d
	var cnt = 1

	for {
		d = d + 1
		if d.String() == dCpy.String() {
			break
		}
		cnt++
	}
	return cnt
}
