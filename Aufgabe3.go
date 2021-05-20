package main

import "fmt"

func main() {
	var testing = false

	if testing {
		proof()
	} else {
		var d = Direction(0)
		var c = Colour(0)
		var cnt = directions()

		t := []TrafficLight{}
		for i := 0; i < cnt; i++ {
			t = append(t, TrafficLight{
				dir: d,
				col: c,
				ax:  d.whichAxis()})
			d = d.next()
			fmt.Println(t[i].String())
		}

		// logik zum schalten
	}
}
