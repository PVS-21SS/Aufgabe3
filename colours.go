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