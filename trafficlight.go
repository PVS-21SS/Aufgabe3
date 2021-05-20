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
	var currentAx = startAx
	for {
		if currentAx == t.ax {
			select {
			case tcol := <-axChan:
				fmt.Println(t.String())
				t.col = tcol

			default:
				fmt.Println(t.String())
				t.col = t.col.next()
				axChan <- t.col
			}
			if t.col == Colour(0) {
				currentAx = currentAx.next()
				select {
				case msgAx := <-allAxChan:
					currentAx = msgAx
				default:
					allAxChan <- currentAx
				}
			}
		}
	}
}
