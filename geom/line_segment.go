package geom

type LineSegment [2]Point

func (segment LineSegment) Endpoint1() Point {
	return segment[0]
}
func (segment LineSegment) Endpoint2() Point {
	return segment[1]
}
