package main

type Colour int

const (
	Red Colour = iota
	Yellow
	Green
)


func (c Colour) String() string {
	col := []string{"Red", "Yellow", "Green"}

	return col[c]
}

func (c Colour) printInColour(light TrafficLight) string{
	col := []string{"Red", "Yellow", "Green"}
	colour := []string{"\033[31m", "\033[32m", "\033[33m"}

	var retString = string(colour[c]) + light.String() + col[c]

	return retString
}

func (c Colour) next() Colour {
	if c < 2 {
		return c+1
	}
	return 0
}

func colourCounter() int {
	var c = Colour(0)
	var ccpy = c
	var cnt = 1
	for {
		c = c.next()
		if c.String() == ccpy.String() {
			break
		}
		cnt++
	}
	return cnt
}