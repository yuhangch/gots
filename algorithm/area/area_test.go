package area

import (
	"fmt"
	"github.com/yuhangch/gots/geom"
	"testing"
)

func TestOfRing(t *testing.T) {
	l, _ := geom.NewLinearRing([]geom.Coordinate{
		{3, 4},
		{5, 11},
		{12, 8},
		{9, 5},
		{5, 6},
		{3, 4},
	})
	fmt.Println(OfRing(l))
}
