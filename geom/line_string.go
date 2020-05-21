package geom

type LineString []Point

func (l LineString) Type() string {
	return "LineString"
}

func (l LineString) Envelope() Envelope {
	return Envelope{}
}
