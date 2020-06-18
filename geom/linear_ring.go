package geom

import "fmt"

// LinearRing struct
type LinearRing []Point

func (r LinearRing) Type() string {
	return "LinearRing"
}

func (r LinearRing) Envelope() *Envelope {
	return pointsEnvelope(r)
}

func (r LinearRing) Valid() bool {
	if len(r) < 3 {
		return false
	}
	if r[0].Equals(r[len(r)-1]) {
		return true
	}
	return false
}

func NewLinearRing(cs []Coordinate) (LinearRing, error) {
	var ring LinearRing
	ring = cs

	if len(ring) < 3 || ring == nil {
		return nil, fmt.Errorf("数据错误")
	}
	if ring.Valid() {
		return ring, nil
	} else {
		return nil, fmt.Errorf("首尾未相接")
	}
}
