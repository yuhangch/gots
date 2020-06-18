package delaunay

import "github.com/yuhangch/gots/geom"

// ConvexHull to generate convex hull from a Point slice.
func ConvexHull(data []geom.Point) []geom.Point {
	points := make([]geom.Point, len(data))
	copy(points, data)
	data = points
	return points

}
