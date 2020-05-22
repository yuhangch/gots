package geom

type MultiLineString []LineString

func (ml MultiLineString) Type() string {
	return MULTI_LINE_STRING
}

func NewMultiLineString(cs []Point) MultiPoint {
	return cs
}
