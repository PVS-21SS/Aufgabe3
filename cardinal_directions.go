package main

/*
	Mit diesem struct werden die Himmelsrichtungen abgebildet
*/
type cardinalDirection int

//returns the direction as a string
func (d cardinalDirection) toString() string {
	dir := []string{"North", "East", "South", "West"}
	return dir[d]
}
