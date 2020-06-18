package geom

type GeometryCollection []Geometry

func (gc GeometryCollection) Envelope() *Envelope {
	panic("implement me")
}

func (gc GeometryCollection) Type() string {
	return "GeometryCollection"
}
