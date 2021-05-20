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

	axChan := make(chan Colour)
	allAxChan := make(chan Axis, 1)
	//var axChan = make(chan Colour)

	t := []TrafficLight{}
	for i := 0; i < cnt; i++ {
		t = append(t, TrafficLight{
			dir: d,
			col: c,
			ax:  d.whichAxis()})
		d = d.next()
		fmt.Println(t[i].String())
	}
	fmt.Println()

	allAxChan <- Axis(0)

	for i := 0; i < cnt; i++ {
		go t[i].run(axChan, allAxChan)
	}
	time.Sleep(time.Millisecond * 10)
}
