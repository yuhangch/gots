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
	txt := "LINESTRING (30 10, 10 30, 40 40)"
	l := NewLexer(txt)
	p := NewParser(l)
	p.parseGeometryType()
}
