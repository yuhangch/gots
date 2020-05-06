package simplify

import (
	"github.com/yuhangch/gots/geom"
	"testing"
)

func TestDp(t *testing.T) {
	line := geom.LineString{
		geom.Point{X: -10.0, Y: 0.0},
		geom.Point{X: -8, Y: 2},
		geom.Point{X: -7, Y: 1},
		geom.Point{X: -3, Y: 4},
		geom.Point{X: -2, Y: 2},
		geom.Point{X: 2, Y: -1},
		geom.Point{X: 5, Y: 0},
		geom.Point{X: 6, Y: 1},
		geom.Point{X: 7, Y: 2},
		geom.Point{X: 10, Y: 3}}

	ans := dp(line, 1)
	if len(ans) != 4 {
		t.Errorf("error for dp test")
	}
}
