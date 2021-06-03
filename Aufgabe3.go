package main

import (
	"time"
)

func main() {
	//create the needed Channels, one for each axis and one for all
	nsAxisChan := make(chan CardinalDirection)
	ewAxisChan := make(chan CardinalDirection)
	wait := make(chan CardinalDirection)

	//create the needed axis, with each two directions and a channel for axis communication
	nsAxis := Axis{dirA: NORTH, dirB: SOUTH, Channel: nsAxisChan}
	ewAxis := Axis{dirA: EAST, dirB: WEST, Channel: ewAxisChan}

	//create the four trafficlights, each with an axis and one direction
	northTraffic := TrafficLight{dir: NORTH, ax: nsAxis}
	eastTraffic := TrafficLight{dir: EAST, ax: ewAxis}
	southTraffic := TrafficLight{dir: SOUTH, ax: nsAxis}
	westTraffic := TrafficLight{dir: WEST, ax: ewAxis}

	//define start axis
	startAxis := nsAxis

	//start the run method of a trafficlight as a goroutine
	go northTraffic.run(startAxis, wait)
	go eastTraffic.run(startAxis, wait)
	go southTraffic.run(startAxis, wait)
	go westTraffic.run(startAxis, wait)

	//give the programm time to run
	time.Sleep(1 * time.Millisecond)
}
