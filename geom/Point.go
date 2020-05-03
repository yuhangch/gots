package geom

import "math"

type Point Coordinate

func (point *Point) Distance(target Point) float64 {
	deltaX, deltaY := point.X-target.X, point.Y-target.Y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
