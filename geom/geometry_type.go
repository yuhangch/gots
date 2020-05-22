package geom

type GeometryType string

const (
	POINT               GeometryType = "Point"
	LINE_STRING                      = "LineString"
	LINEAR_RING                      = "LinearRing"
	POLYGON                          = "Polygon"
	MULTI_POINT                      = "MultiPoint"
	MULTI_LINE_STRING                = "MultiLineString"
	MULTI_POLYGON                    = "MultiPolygon"
	GEOMETRY_COLLECTION              = "GeometryCollection"
	CIRCLE                           = "Circle"
)
