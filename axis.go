package main

type Axis int
const (
	nsAxis = iota
	ewAxis
)

func (axis Axis) String() string {
	return [...]string{"nsAxis", "ewAxis"}[axis]
}
