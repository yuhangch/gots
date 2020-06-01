package strtree

import (
	"github.com/yuhangch/gots/geom"
	"sort"
)

type Bounds []*geom.Envelope

func (b *Bounds) Len() int {
	return len(*b)
}

func (b *Bounds) Less(i, j int) bool {
	ix, _ := (*b)[i].CentreXY()
	jx, _ := (*b)[j].CentreXY()
	return ix < jx
}

func (b *Bounds) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *Bounds) Sort() {
	sort.Sort(b)
}
