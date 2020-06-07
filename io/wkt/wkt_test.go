package wkt

import (
	"testing"
)

func TestParsePoint(t *testing.T) {
	txt := "POINT (-30.3333 +10.1111)"
	g := Read(txt)
	Write(g)
}

func TestParseLineString(t *testing.T) {
	txt := "LINESTRING (30.6666 10, 10.9999 30, 40 40)"
	g := Read(txt)
	Write(g)
}

func TestParsePolygon(t *testing.T) {
	txt := "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"
	g := Read(txt)
	Write(g)
}

func TestParseMultiPoint(t *testing.T) {
	txt := "MULTIPOINT ((10 40), (40 30), (20 20), (30 10))"
	Read(txt)

	txt2 := "MULTIPOINT (10 40, 40 30, 20 20, 30 10)"
	Read(txt2)
}

func TestParseMultiLineString(t *testing.T) {
	txt := "MULTILINESTRING ((10 10, 20 20, 10 40), (40 40, 30 30, 40 20, 30 10))"
	Read(txt)
}

func TestParseMultiPolygon(t *testing.T) {
	txt := "MULTIPOLYGON (((30 20, 45 40, 10 40, 30 20)), ((15 5, 40 10, 10 20, 5 10, 15 5)))"
	Read(txt)
	txt2 := "MULTIPOLYGON (((40 40, 20 45, 45 30, 40 40)), ((20 35, 10 30, 10 10, 30 5, 45 20, 20 35), (30 20, 20 15, 20 25, 30 20)))"
	Read(txt2)
}

func TestParseGeometryCollection(t *testing.T) {
	txt := "GEOMETRYCOLLECTION (POINT (40 10), LINESTRING (10 10, 20 20, 10 40), POLYGON ((40 40, 20 45, 45 30, 40 40)))"
	Read(txt)

}
