package geom

type MultiPolygon []Polygon

func (mp MultiPolygon) Type() string {
	return MULTI_POLYGON
}
