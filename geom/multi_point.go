package geom

type MultiPoint []Point

func (mp MultiPoint) Type() string {
	return MULTI_POINT
}

func (mp MultiPoint) Envelope() *Envelope {
	return pointsEnvelope(mp)
}

func NewMultiPoint(cs []Point) MultiPoint {
	return cs
}
