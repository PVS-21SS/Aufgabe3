//go build -race
package main

import (
	"fmt"
	"time"
)

func main() {
	var d = Direction(0)
	var c = Colour(0)
	var cnt = directionCounter()


	// Channel for the Colour, to sync the to TrafficLights of an Axis
	axChanColour := make(chan Colour)
	// Channel for the Axis, the Axis has two TrafficLights
	// the two TrafficLights could be: North and South, with these Directions
	axisDirectionChan := make(chan Axis, 1)
	// quitChannel to Stop the Running threads
	quitChannel := make(chan bool)

	// initialisation of the TrafficLights
	var t []TrafficLight
	// cnt is the Counter for the directions
	for i := 0; i < cnt; i++ {
		// every TrafficLight, gets a Direction, a Colour and a Axis
		t = append(t, TrafficLight{
			dir: d,
			col: c,
			ax:  d.whichAxis()})
		d = d.next()
		// initial print of the TrafficLights
		fmt.Println(t[i].printInColour())
	}
	fmt.Println()

	// setting the Starting Axis, 0 means North and South
	// 1 would be East and West
	axisDirectionChan <- Axis(0)

	// starting all trafficLights in their own Thread
	for i := 0; i < cnt; i++ {
		go t[i].run(axChanColour, axisDirectionChan, quitChannel)
	}
	// wait a while to Show the TrafficLights are working
	time.Sleep(time.Millisecond * 10)
	// Stop all goroutines
	quitChannel <- true
	// wait fot the goroutines to be done
	<- quitChannel

}
