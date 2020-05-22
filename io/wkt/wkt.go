package wkt

import "github.com/yuhangch/gots/geom"

func Read(wkt string) geom.Geometry {
	p := NewParser(wkt)
	return p.parse()
}

func Write(g geom.Geometry) string {
	return encode(g)
}
