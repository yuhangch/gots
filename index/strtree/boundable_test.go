package strtree

import (
	"fmt"
	"github.com/yuhangch/gots/geom"
	"testing"
)

func TestSort(t *testing.T) {
	var b Bounds
	e1 := geom.NewEnvelope([4]float64{0, 1, 1, 0})
	e2 := geom.NewEnvelope([4]float64{1, 2, 2, 1})
	e3 := geom.NewEnvelope([4]float64{2, 3, 3, 2})
	e4 := geom.NewEnvelope([4]float64{3, 4, 4, 3})
	b = append(b, e2, e1, e4, e3)
	b.Sort()
	for _, v := range b {
		fmt.Println(v.Centre())

	}

}
