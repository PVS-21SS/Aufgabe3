package main

import "fmt"

type TrafficLight struct {
	dir CardinalDirection
	ax  Axis
	col Colour
}

func (t TrafficLight) printInColour() string {
	colour := []string{"\033[31m", "\033[32m", "\033[33m"}
	return colour[t.col] + t.dir.String() + ":\t" + t.col.String()
}

func (t *TrafficLight) run(startAxis Axis, waitingAxis chan CardinalDirection) {

	if t.ax != startAxis {
		<-waitingAxis
	}

	for {
		fmt.Println(t.printInColour())

		if t.col == 0 {
			t.synchronize()

			waitingAxis <- t.dir
			t.synchronize()

			<-waitingAxis
			t.col.nextColour()
		} else {

			t.synchronize()

			t.col.nextColour()
		}
	}
}

func (t *TrafficLight) synchronize() {
	select {
	case <-t.ax.Channel:
	case t.ax.Channel <- t.dir:
	}
}
