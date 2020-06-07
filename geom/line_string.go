package geom

type LineString []Point

func (l LineString) Type() string {
	return LINE_STRING
}

func (l LineString) Envelope() *Envelope {
	return pointsEnvelope(l)
}

func NewLineString(cs []Point) LineString {
	return cs

}

func pointsEnvelope(pts []Point) *Envelope {
	e := NewEmptyEnvelope()
	for _, v := range pts {
		e.Expand(v)
	}
	return e
}
