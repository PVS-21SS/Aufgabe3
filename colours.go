package main


type Colour int

const (
	RED    Colour = iota
	GREEN
	YELLOW
)

func (c *Colour) nextColour() {
	switch *c {
	case RED:
		*c = GREEN
	case GREEN:
		*c = YELLOW
	case YELLOW:
		*c = RED
	}
}

func (c Colour) String() string {
	col := []string{"Red", "Green", "Yellow"}
	return col[c]
}
