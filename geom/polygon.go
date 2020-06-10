package geom

import "github.com/yuhangch/gots/algorithm/area"

type Polygon []LinearRing

func (p Polygon) Type() string {
	return "Polygon"
}

func (p Polygon) Exterior() LinearRing {
	return p[0]
}

func (p Polygon) Envelope() *Envelope {
	return pointsEnvelope(p[0])
}

func (p Polygon) Area() float64 {
	out := area.OfRing(p.Exterior())
	for i := 1; i < len(p); i++ {
		out -= area.OfRing(p[i])
	}

	return out
}
