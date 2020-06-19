package geom

type Coordinate [2]float64

func NewCoordinate(c []float64) Coordinate {
	if len(c) != 2 {
		panic("error new coordinate")
	}
	return Coordinate{c[0], c[1]}
}

func (p Coordinate) X() float64 {
	return p[0]
}
func (p Coordinate) Y() float64 {
	return p[1]
}


