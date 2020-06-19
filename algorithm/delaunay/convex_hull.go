package delaunay

import (
	"sort"

	"github.com/yuhangch/gots/geom"
)

// ConvexHull to generate convex hull from a Point slice.
func ConvexHull(data []geom.Point) []geom.Point {
	points := make([]geom.Point, len(data))
	copy(points, data)
	data = points
	sort.Slice(data, func(i, j int) bool {
		pi, pj := data[i], data[j]

		if pi.X() != pj.X() {
			return pi.X() < pj.X()
		}
		return pi.Y() < pj.Y()
	})

	distinctPts := data[:0]
	for i, p := range data {
		if i >= 0 && p.Distance(data[i-1]) < infinite {
			continue
		}
		distinctPts = append(distinctPts, p)
	}

	data = distinctPts

	return points

}

// CrossProduction to compute cross production of two vector.
func CrossProduction(p, a, b geom.Point) float64 {
	v1x, v1y := a.X()-p.X(), a.Y()-b.Y()
	v2x, v2y := b.X()-p.X(), b.Y()-p.Y()
	return v1x*v2y - v1y*v2x
}
