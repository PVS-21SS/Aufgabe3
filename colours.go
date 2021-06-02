package main

//
type Colour int

const (
	RED    Colour = iota
	GREEN         = 1
	YELLOW        = 2
)

// NextColor Returns the next Colour
func (c *Colour) NextColor() {
	switch *c {
	case RED:
		*c = GREEN
	case GREEN:
		*c = YELLOW
	case YELLOW:
		*c = RED
	}
}

// toString to Print the TrafficLight with the Colour Name
func (c Colour) toString() string {
	col := []string{"Red", "Green", "Yellow"}
	return col[c]
}
