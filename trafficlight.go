package main

import "fmt"

type TrafficLight struct {
	dir CardinalDirection
	ax  Axis
	col Colour
}

//returns a string that conatins the direction of the trafficlight and the colour
//the String is coloured based on the colour of the trafficlight
func (t TrafficLight) printInColour() string {
	colour := []string{"\033[31m", "\033[32m", "\033[33m"}
	return colour[t.col] + t.dir.String() + ":\t" + t.col.String()
}

//the run method gets the startAxis and a channel for syncronising
func (t *TrafficLight) run(startAxis Axis, waitingAxis chan CardinalDirection) {

	//if the trafficlight is not on the startAxis it has to wait
	if t.ax != startAxis {
		<-waitingAxis
	}

	//infinite loop to run the trafficlight until it is stopped
	for {
		fmt.Println(t.printInColour())

		// if Color is Red, change Axis and Colour to next Colour
		if t.col == RED {
			t.synchronize()

			waitingAxis <- t.dir
			t.synchronize()

			<-waitingAxis
			t.col.nextColour()
		} else {
			// change Colour to the Point its Red
			t.synchronize()

			t.col.nextColour()
		}
	}
}

// Sync the Axis over the Axis Channel
func (t *TrafficLight) synchronize() {
	select {
	case <-t.ax.Channel:
	case t.ax.Channel <- t.dir:
	}
}
