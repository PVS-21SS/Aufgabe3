package main

import "fmt"

type TrafficLight struct {
	dir Direction
	col Colour
	ax  Axis
}

func (t TrafficLight) String() string {
	return t.dir.String() + ": " + t.col.String()
}

func (t TrafficLight) run(startAx Axis) {

	for i := 0; i < 3; i++ {
		if t.ax == startAx {
			t.col = t.col.next()
			fmt.Println(t.String())
		}
	}
}
