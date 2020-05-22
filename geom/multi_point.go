package geom

type MultiPoint []Point

func (mp MultiPoint) Type() string {
	return MULTI_POINT
}

func NewMultiPoint(cs []Point) MultiPoint {
	return cs
}
