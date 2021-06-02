package main

import "fmt"

type TrafficLight struct {
	dir cardinalDirection
	ax  Axis
	col Colour
}

// ShowColor Prints the direction and colour of the trafficlight
func (t TrafficLight) ShowColor() string {
	colour := []string{"\033[31m", "\033[32m", "\033[33m"}

	return colour[t.col] + t.dir.toString() + ":\t" + t.col.toString()
}

func (t *TrafficLight) run(startAxis Axis, waitingAxis chan cardinalDirection) {

	// if the axis of a trafficlight doesn't is not the start axis, wait for signal
	if t.ax != startAxis {
		_ = <-waitingAxis
	}

	for {
		fmt.Println(t.ShowColor())

		// if the trafficlight colour is red switch to next axis
		if t.col == 0 {
			t.syncronise()

			waitingAxis <- t.dir
			t.syncronise()

			_ = <-waitingAxis
			t.col.NextColor()
		} else {
			// Die Ampeln synchronisieren sich bevor sie zur nächsten Farbe schalten
			//
			t.syncronise()

			// Gibt die Farbe der Ampel aus und schaltet zur nächsten Farbe weiter
			t.col.NextColor()
		}
	}
}

/*
	Diese Funktion synchronisiert zwei TrafficLights
*/
func (t *TrafficLight) syncronise() {
	select {
	case _ = <-t.ax.Channel:
	case t.ax.Channel <- t.dir:
	}
}
