package wkt

import (
	"fmt"
	"github.com/yuhangch/gots/geom"
	"strings"
)

func ptsToStr(points []geom.Point) string {
	var merge []string
	if len(points) < 1 {
		return ""
	}
	for _, v := range points {
		merge = append(merge, fmt.Sprintf("%f %f", v.X(), v.Y()))
	}
	return strings.Join(merge, ", ")

}

func encodePoint(p geom.Point) string {
	return fmt.Sprintf("%f %f", p.X(), p.Y())

}

func encodeLineString(l geom.LineString) string {
	return ptsToStr(l)
}

func encode(g geom.Geometry) string {
	var enc string
	geomType := g.Type()
	switch geomType {
	case "Point":
		enc = encodePoint(g.(geom.Point))
	case "LineString":
		enc = encodeLineString(g.(geom.LineString))

	}
	wkt := fmt.Sprintf("%s (%s)", strings.ToUpper(geomType), enc)
	fmt.Println(wkt)
	return wkt
}
