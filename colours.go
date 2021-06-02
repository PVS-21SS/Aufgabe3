package main

//Definition Colour als enum
type Colour int

const (
	red Colour = iota
	green
	yellow
)

func nextColour(colour Colour) Colour {
	return (colour + 1) % (yellow + 1)
}

func (colour Colour) String() string {
	return [...]string{"Red", "Green", "Yellow"}[colour]
}