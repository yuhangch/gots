package strtree

import (
	"github.com/yuhangch/gots/geom"
	"sort"
)

type Bounds []*geom.Envelope

func (b *Bounds) Len() int {
	return len(*b)
}

func (b *Bounds) LessX(i, j int) bool {
	ix, _ := (*b)[i].CentreXY()
	jx, _ := (*b)[j].CentreXY()
	return ix < jx
}

func (b *Bounds) LessY(i, j int) bool {
	_, iy := (*b)[i].CentreXY()
	_, jy := (*b)[j].CentreXY()
	return iy < jy
}

func (b *Bounds) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *Bounds) SortX() {
	sort.SliceStable(*b, b.LessX)
}

func (b *Bounds) SortY() {
	sort.SliceStable(*b, b.LessY)
}
