package main

import "fmt"

func proof(){
	var d = Direction(0)
	var c = Colour(0)
	var cnt = directionCounter()

	fmt.Println("First Direction: \t\t", d)
	fmt.Println("Next direction: \t\t", d.next())
	fmt.Println( "opposite direction: \t", d.opposite())
	fmt.Println("Axis: \t\t\t\t\t", d.whichAxis(), )

	fmt.Println()

	fmt.Println("First Colour:\t\t\t", c)
	fmt.Println("Next Colour:\t\t\t", c.next())


	fmt.Println("Count directionCounter:\t", cnt)

	fmt.Println()
}
