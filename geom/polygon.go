package geom

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
