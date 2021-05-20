package main

import "fmt"

type TrafficLight struct {
	dir Direction
	col Colour
	ax  Axis
}

func (t TrafficLight) String() string {
	var s = t.dir.String() + ": " + t.col.String()
	return s
}
func (t TrafficLight) run() {
	for {
		fmt.Println(t.String())

		t.col = t.col.next()
	}
}
