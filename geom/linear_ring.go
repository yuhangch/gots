package geom

import "fmt"

type LinearRing LineString

func (r LinearRing) Type() string {
	return "LineString"
}

func (r LinearRing) Envelope() Envelope {
	return Envelope{}
}

func (r LinearRing) Valid() bool {
	if len(r) == 0 {
		return false
	}
	if r[0].Equals(r[len(r)-1]) {
		return true
	}
	return false
}

func NewLinearRing(cs []Coordinate) (LinearRing, error) {
	var ring LinearRing
	for _, v := range cs {
		ring = append(ring, v)
	}
	if len(ring) < 1 || ring == nil {
		return nil, fmt.Errorf("数据为空")
	}
	if ring.Valid() {
		return ring, nil
	} else {
		return nil, fmt.Errorf("首尾未相接")
	}
}
