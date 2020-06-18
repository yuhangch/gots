package geom

type MultiPoint []Point

func (mp MultiPoint) Type() string {
	return "MultiPoint"
}

func (mp MultiPoint) Envelope() *Envelope {
	return pointsEnvelope(mp)
}

func NewMultiPoint(cs []Point) MultiPoint {
	return cs
}
