package simplify

import (
	"github.com/yuhangch/gots/geom"
	"testing"
)

func TestDp(t *testing.T) {
	line := geom.LineString{
		geom.Point{-10.0, 0.0},
		geom.Point{-8, 2},
		geom.Point{-7, 1},
		geom.Point{-3, 4},
		geom.Point{-2, 2},
		geom.Point{2, -1},
		geom.Point{5, 0},
		geom.Point{6, 1},
		geom.Point{7, 2},
		geom.Point{10, 3}}

	ans := DP(line, 1)
	if len(ans) != 4 {
		t.Errorf("error for DP test")
	}
}
