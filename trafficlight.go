package main

import (
	"fmt"
)

// TrafficLight
/*
 every Trafficlight gets 3 Parameters
 The Direction of the Trafficlight
 The Colour of the Trafficlight(Only Parameter that will get changed)
 The Axis, two Trafficlights build an Axis
*/

var nsChannel = make (chan Axis)
var ewChannel = make (chan Axis)
var sync = make (chan Colour)

func trafficlight(direction CardinalDirection, quitChannel chan bool) {

	var trafficlightColour = Colour(green)
	show(trafficlightColour, direction)
	for {
		show(trafficlightColour, direction)
		select {
		default:
			// if the Colour is Red, write the next Axis to the currentAxis Variable
			if trafficlightColour == red {

				ewChannel <- ewAxis
				switchAxis(nsAxis)
			}
			trafficlightColour = nextColour(trafficlightColour)

		case <-quitChannel:
			quitChannel <- true
			return
		}

	}
}

func show(colour Colour, direction CardinalDirection) {
	colours := map[Colour]string{
		red : "\033[31m",
		green : "\033[32m",
		yellow : "\033[33m"}

	fmt.Println(colours[colour], direction, ":" , colour)
}

func getAxisChannel(direction CardinalDirection) chan Axis {
	if direction == north || direction == east {
		return nsChannel
	}
	return ewChannel
}

func switchAxis(direction CardinalDirection) chan Axis{
	if direction == north || direction == south {
		return nsChannel
	}
	return ewChannel
}


