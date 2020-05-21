package geom

type Polygon []LinearRing

func (p Polygon) Type() string {
	return "Polygon"
}

func (p Polygon) Envelope() Envelope {
	return Envelope{}
}
