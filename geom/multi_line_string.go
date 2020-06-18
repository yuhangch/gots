package geom

type MultiLineString []LineString

func (ml MultiLineString) Type() string {
	return "MultiLineString"
}

func (ml MultiLineString) Envelope() *Envelope {
	e := NewEmptyEnvelope()
	for _, v := range ml {
		e.Expand(v.Envelope())
	}
	return e
}

func NewMultiLineString(cs []Point) MultiPoint {
	return cs
}
