package index

import "github.com/yuhangch/gots/geom"

type SpatialIndex interface {
	Insert(envelope geom.Envelope, item interface{})
	Query(envelope geom.Envelope)
	QueryAndApply(envelope geom.Envelope, item interface{})
	Remove(envelope geom.Envelope, item interface{})
}
