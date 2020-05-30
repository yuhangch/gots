package geom

import (
	"fmt"
	"testing"
)

func TestEnvelope_Contains(t *testing.T) {
	e := NewEnvelope([4]float64{0, 4, 4, 0})
	p := Point([2]float64{2, 9})
	fmt.Println(e.Contains(&p))
	o := NewEnvelope([4]float64{1, 2, 2, 1})
	fmt.Println(e.Contains(&o))

}
