package main

/*
	Mit diesem struct werden die Achsen definiert
*/
type Axis struct {
	dirA    cardinalDirection
	dirB    cardinalDirection
	Channel chan cardinalDirection // Der Kommunikations Channel für diese Achse
}
/*
	Gibt einen String auf der Basis der beiden cardinalDirections zurück
*/
func (a Axis) ToString() string {
	return a.dirA.toString() + a.dirB.toString()
}
