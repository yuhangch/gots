package geom

type GeometryCollection []Geometry

func (gc GeometryCollection) Type() string {
	return GEOMETRY_COLLECTION
}
