package geom

type MultiPolygon []Polygon

func (mp MultiPolygon) Envelope() *Envelope {
	e := NewEmptyEnvelope()
	for _, v := range mp {
		e.Expand(v.Exterior())
	}
	return e
}

func (mp MultiPolygon) Type() string {
	return MULTI_POLYGON
}
