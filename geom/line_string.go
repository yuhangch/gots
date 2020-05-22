package geom

type LineString []Point

func (l LineString) Type() string {
	return LINE_STRING
}

func (l LineString) Envelope() Envelope {
	return Envelope{}
}

func NewLineString(cs []Point) LineString {
	return cs
}
