package main

import (
	"time"
)

func main() {
	// Erzeugen des Kommunikation Channels
	NS := make(chan cardinalDirection)
	EW := make(chan cardinalDirection)
	wait := make(chan cardinalDirection)

	// Erzeugen der Himmelsrichtungen
	north := cardinalDirection(0)
	east := cardinalDirection(1)
	south := cardinalDirection(2)
	west := cardinalDirection(3)

	// Erzeugen der Achsen
	nsAxis := Axis{dirA: north, dirB: south, Channel: NS}
	ewAxis := Axis{dirA: east, dirB: west, Channel: EW}

	// Definition der StartAxis
	startAxis := nsAxis

	// In diesem Fall rot
	startColor := Colour(0)

	// Erzeugen der TrafficLights
	northTraffic := TrafficLight{dir: north, ax: nsAxis, col: startColor}
	eastTraffic := TrafficLight{dir: east, ax: ewAxis, col: startColor}
	southTraffic := TrafficLight{dir: south, ax: nsAxis, col: startColor}
	westTraffic := TrafficLight{dir: west, ax: ewAxis, col: startColor}

	// Erzeugen der Goroutinen
	go northTraffic.run( startAxis, wait)
	go eastTraffic.run( startAxis, wait)
	go southTraffic.run( startAxis, wait)
	go westTraffic.run( startAxis, wait)

	// Wartet damit die Goroutinen nicht abgebrochen haben
	time.Sleep(1 * time.Millisecond)
}
