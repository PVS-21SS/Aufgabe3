package main

type Direction int

func (d Direction) String() string {
	dir := []string{"North", "East", "South", "West"}
	if int(d) >= cap(dir) {
		d = Direction(int(d) - cap(dir))
	}
	return dir[d]
}

func (d Direction) next() Direction {
	if d < Direction(directionCounter()-1) {
		return d + 1
	}
	return 0
}

func (d Direction) opposite() Direction {
	return d + Direction(directionCounter()/2)
}
func (d Direction) whichAxis() Axis {
	return Axis(d % 2)
}

func directionCounter() int {
	var d = Direction(0)
	var dcpy = d
	var cnt = 1
	for {
		d = d+1
		if d.String() == dcpy.String() {
			break
		}
		cnt++
	}
	return cnt
}
