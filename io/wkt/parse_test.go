package wkt

import (
	"testing"
)

func TestParsePoint(t *testing.T) {
	txt := "POINT (-30.3333 +10.1111)"
	l := NewLexer(txt)
	p := NewParser(l)
	p.parseGeometryType()
}

func TestParseLineString(t *testing.T) {
	txt := "LINESTRING (30.6666 10, 10.9999 30, 40 40)"
	l := NewLexer(txt)
	p := NewParser(l)
	p.parseGeometryType()
}

func TestParsePolygon(t *testing.T) {
	txt := "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"
	l := NewLexer(txt)
	p := NewParser(l)
	p.parseGeometryType()
}
