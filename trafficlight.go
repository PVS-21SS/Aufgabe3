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

func (t TrafficLight) run(axChanColour chan Colour, axisDirectionChan chan Axis) {
	for {
		switch currentAx := <-axisDirectionChan; {
		case currentAx == t.ax:
			fmt.Println(t.col.printInColour(t))
			select {

			case tcol := <-axChanColour:
				//fmt.Println(t.String())
				t.col = tcol
				if t.col == Colour(0) {
					currentAx = currentAx.next()
				}
				axisDirectionChan <- currentAx

			default:
				t.col = t.col.next()
				axisDirectionChan <- currentAx
				axChanColour <- t.col
				//fmt.Println(t.String())
			}

		default:
			axisDirectionChan <- currentAx
		}

	}
}
