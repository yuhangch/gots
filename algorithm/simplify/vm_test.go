package simplify

import (
	"fmt"
	"github.com/yuhangch/gots/geom"
	"testing"
)

func TestVm(t *testing.T) {
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

	ans := VM(line, 3)
	//if len(ans) != 4 {
	//	t.Errorf("error for DP test")
	//}
	fmt.Println(len(ans), ans)
}
