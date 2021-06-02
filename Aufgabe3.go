package main

import (
	"fmt"
	"time"
)

// helper Function
// to print coloured Text and try to use Maps
func printColouredMessage(colour string, msg string) {
	colours := map[string]string{
		"Red" : "\033[31m",
		"Yellow" : "\033[32m",
		"Green" : "\033[33m",
		"Blue" : "\033[34m",
		"Purple" : "\033[35m",
		"Cyan" : "\033[36m",
		"White" : "\033[37m",
		"Reset" : "\033[0m",
	}
	fmt.Println(colours[colour] + msg + "\033[0m")
}


func main() {
	// Channel for the Colour, to sync the to TrafficLights of an Axis
	//axChanColour := make(chan Colour)
	// Channel for the Axis, the Axis has two TrafficLights
	// the two TrafficLights could be: North and South, with these Directions
	//axisDirectionChan := make(chan Axis, 1)
	// quitChannel to Stop the Running threads
	quitChannel := make(chan bool)

	printColouredMessage("Blue", "\nStarting the Intersection!\n")

	// starting all trafficLights in their own Thread
	go trafficlight(north, quitChannel)
	go trafficlight(south, quitChannel)
	go trafficlight(east, quitChannel)
	go trafficlight(south, quitChannel)

	// wait a while to Show the TrafficLights are working
	time.Sleep(time.Millisecond * 1)

	// Stop all goroutines
	quitChannel <- true

	// wait fot the goroutines to be done
	<- quitChannel
	printColouredMessage("Blue", "Intersection Stopped!")
}
