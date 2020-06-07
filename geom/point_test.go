package geom

import (
	"fmt"
	"testing"
)

func TestPoint_Envelope(t *testing.T) {
	p := Point{1, 2}
	fmt.Println(p.Envelope())
}

func TestPoint_Distance(t *testing.T) {

}
