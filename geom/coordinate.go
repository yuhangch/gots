package geom

type Coordinate [2]float64

func NewCoordinate(a, b float64) Coordinate {
	return Coordinate([2]float64{a, b})
}

func (p Coordinate) X() float64 {
	return p[0]
}
func (p Coordinate) Y() float64 {
	return p[1]
}
