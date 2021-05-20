package main

import (
	"fmt"
)

type TrafficLight struct {
	dir Direction
	col Colour
	ax  Axis
}

func (t TrafficLight) String() string {
	return t.dir.String() + ": " + t.col.String()
}

func (t TrafficLight) run(startAx Axis, axChan chan Colour, allAxChan chan Axis) {
	var cnt = colourCounter()
	for {
		if startAx == t.ax {
			select {
			case tcol := <-axChan:
				fmt.Println(t.String())
				fmt.Println("----------------------")
				t.col = tcol
				if tcol == Colour(cnt) {
				}
			default:
				fmt.Println(t.String())
				t.col = t.col.next()
				axChan <- t.col
			}
		}
	}
}
