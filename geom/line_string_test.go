package geom

import (
	"fmt"
	"testing"
)

func TestLineString(t *testing.T) {
	line := LineString{Point{1, 2}, Point{2, 3}}
	fmt.Println(len(line))
}

func TestLineString_Envelope(t *testing.T) {
	line := LineString{Point{1, 1}, Point{2, 2}, Point{3, 3}}
	fmt.Println(line.Envelope())
}
