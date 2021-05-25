package main

import "fmt"

// TrafficLight
/*
 every Trafficlight gets 3 Parameters
 The Direction of the Trafficlight
 The Colour of the Trafficlight(Only Parameter that will get changed)
 The Axis, two Trafficlights build an Axis
*/

type TrafficLight struct {
	dir Direction
	col Colour
	ax  Axis
}

// toString function for the TrafficLight
func (t TrafficLight) String() string {
	return t.dir.String() + ": " + t.col.String()
}

// print TrafficLight and Colour in ANSI Colours
func (t TrafficLight) printInColour() string {
	colour := []string{"\033[31m", "\033[33m", "\033[32m"}

	var retString = colour[t.col] + t.dir.String() + ": " + t.col.String()

	return retString
}

func (t TrafficLight) run(axChanColour chan Colour, axisDirectionChan chan Axis, quitChannel chan bool) {
	fmt.Println(t.printInColour())
	for {
		select {
		default:
			// get current Axis
			switch currentAx := <-axisDirectionChan; {

			//  Check if TrafficLight is Part of Axis
			case currentAx == t.ax:
				//fmt.Println(t.printInColour())
				select {

				// if axChanColour has a Value, run this Code
				case tcol := <-axChanColour:

					// change the own Colour to the one of axChanColour
					t.col = tcol
					fmt.Println(t.printInColour())
					// if the Colour is Red, write the next Axis to the currentAxis Variable
					if t.col == Colour(0) {
						currentAx = currentAx.next()
					}

					// write the currentAxis to the Channel
					axisDirectionChan <- currentAx

				// if axChan has no Value, write the next Colour to the TrafficLight and to the axChanColour Channel
				default:
					t.col = t.col.next()
					fmt.Println(t.printInColour())
					axisDirectionChan <- currentAx
					axChanColour <- t.col
				}

			default:
				// if TrafficLight is not Part of the Axis, it returns the Axis into the Channel
				axisDirectionChan <- currentAx
			}

		case <-quitChannel:
			quitChannel <- true
			return
		}

	}
}
