package main
type TrafficLight struct {
	dir Direction
	col Colour
	ax Axis
}
func (t TrafficLight) String() string {
	var s = t.dir.String() + ": "+ t.col.String()
	return s
}
