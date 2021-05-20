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

func (t TrafficLight) run(axChan chan Colour, allAxChan chan Axis) {

	for {
		currentAx := <- allAxChan
		allAxChan <- currentAx
		if currentAx == t.ax {
			//fmt.Println(t.ax, t.String())
			fmt.Println(t.String())
			select {

			case tcol := <-axChan:
				fmt.Println("------------")
				t.col = tcol

			default:
				t.col = t.col.next()
				axChan <- t.col
			}

			if t.col == Colour(0) {
				//fmt.Println(t.String())
				currentAx = currentAx.next()
				allAxChan <- currentAx
			}
		}
	}
}
