package geom

import "math"

type Point = Coordinate

func (p Point) Distance(target Point) float64 {
	deltaX, deltaY := p.X()-target.X(), p.Y()-target.Y()
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func (p Point) Type() string {
	return string(PointType)
}

func (p Point) Envelope() Envelope {
	return Envelope{}
}

func (p Point) Equals(p2 Point) bool {
	if p.X() == p2.X() && p.Y() == p2.Y() {
		return true
	}
	return false
}
