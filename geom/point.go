package geom

import "math"

type Point = Coordinate

func (p Point) Distance(target Point) float64 {
	deltaX, deltaY := p.X()-target.X(), p.Y()-target.Y()
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func (p Point) Type() string {
	return "Point"
}

func (p Point) Envelope() Envelope {
	return Envelope{}
}
