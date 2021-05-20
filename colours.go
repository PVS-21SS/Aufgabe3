package main

type Colour int

const (
	Red Colour = 0
	Yellow
	Green
)

func (c Colour) String() string {
	col := []string{"Red", "Yellow", "Green"}
	if int(c) >= cap(col) {
		c = Colour(int(c) - cap(col))
	}
	return col[c]
}
func (c Colour) next() Colour {
	return c + 1
}
