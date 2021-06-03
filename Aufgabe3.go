package main

import (
	"time"
)

func main() {


	nsAxisChan := make(chan CardinalDirection)
	ewAxisChan := make(chan CardinalDirection)
	wait := make(chan CardinalDirection)


	nsAxis := Axis{dirA: north, dirB: south, Channel: nsAxisChan}
	ewAxis := Axis{dirA: east, dirB: west, Channel: ewAxisChan}


	northTraffic := TrafficLight{dir: north, ax: nsAxis}
	eastTraffic := TrafficLight{dir: east, ax: ewAxis}
	southTraffic := TrafficLight{dir: south, ax: nsAxis}
	westTraffic := TrafficLight{dir: west, ax: ewAxis}


	startAxis := nsAxis


	go northTraffic.run(startAxis, wait)
	go eastTraffic.run(startAxis, wait)
	go southTraffic.run(startAxis, wait)
	go westTraffic.run(startAxis, wait)

	time.Sleep(1 * time.Millisecond)
}
