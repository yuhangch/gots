package geom

import (
	"math"
)

type Envelope [4]float64

func (e *Envelope) Left() float64 {
	return e[0]
}

func (e *Envelope) Top() float64 {
	return e[1]
}

func (e *Envelope) Right() float64 {
	return e[2]
}

func (e *Envelope) Bottom() float64 {
	return e[3]
}

func NewEmptyEnvelope() *Envelope {
	return &Envelope{0, -1, -1, 0}
}

func NewEnvelope(coordinates [4]float64) *Envelope {
	var env Envelope = coordinates
	if env.Height() < 0 {
		env[1], env[3] = env[3], env[1]
	}
	if env.Width() < 0 {
		env[0], env[2] = env[2], env[0]
	}
	return &env
}

func (e *Envelope) isEmpty() bool {
	if e == nil ||
		e.Width() < 0 ||
		e.Height() < 0 {
		return true
	}
	return false
}

func (e *Envelope) Width() float64 {
	return e.Right() - e.Left()
}

func (e *Envelope) Height() float64 {
	return e.Top() - e.Bottom()
}

func (e *Envelope) Area() (area float64) {
	return e.Width() * e.Height()
}

func (e *Envelope) Diameter() float64 {
	w, h := e.Width(), e.Height()
	return math.Sqrt(w*w + h*h)
}

func (e *Envelope) Centre() *Coordinate {
	if e.isEmpty() {
		return nil
	}
	x, y := (e.Left()+e.Right())/2, (e.Top()+e.Bottom())/2
	c := NewCoordinate(x, y)
	return &c

}
func (e *Envelope) CentreXY() (float64, float64) {
	if e.isEmpty() {
		return 0, 0
	}
	return (e.Left() + e.Right()) / 2, (e.Top() + e.Bottom()) / 2

}

func (e *Envelope) Intersects(other *Envelope) bool {
	if e.isEmpty() || other.isEmpty() {
		return false
	}
	return !(other.Right() < e.Left() ||
		other.Left() > e.Right() ||
		other.Top() < e.Bottom() ||
		other.Bottom() > e.Top())
}

func (e *Envelope) Disjoint(other *Envelope) bool {
	if e.isEmpty() || other.isEmpty() {
		return true
	}
	return other.Right() < e.Left() ||
		other.Left() > e.Right() ||
		other.Top() < e.Bottom() ||
		other.Bottom() > e.Top()
}

func (e *Envelope) Intersection(envelope *Envelope) *Envelope {
	if e.isEmpty() || envelope.isEmpty() || !e.Intersects(envelope) {
		return NewEmptyEnvelope()
	}

	left := math.Max(envelope.Left(), e.Left())
	right := math.Max(envelope.Right(), e.Right())
	top := math.Min(envelope.Top(), e.Top())
	bottom := math.Min(envelope.Bottom(), e.Bottom())
	return NewEnvelope([4]float64{left, top, right, bottom})
}

func (e *Envelope) wrapBool(item interface{}, func1 func(point Point) bool, func2 func(envelope Envelope) bool) bool {
	v := item
	switch v.(type) {
	case Point:
		return func1(v.(Point))
	case Envelope:
		return func2(v.(Envelope))
	default:
		panic("no support type.")
	}
}
func (e *Envelope) wrapVoid(item interface{}, func1 func(point Point), func2 func(envelope Envelope)) {
	v := item
	switch v.(type) {
	case Point:
		func1(v.(Point))
	case Envelope:
		func2(v.(Envelope))
	default:
		panic("no support type.")
	}
}

func (e *Envelope) Contains(item interface{}) bool {
	//v := *item
	//switch v.(type){
	//case Point:
	//	return e.containsPoint(v.(Point))
	//case Envelope:
	//	return e.containsOther(v.(Envelope))
	//default:
	//	return false
	//}
	return e.wrapBool(item, e.containsPoint, e.containsOther)
}

func (e *Envelope) containsPoint(p Point) bool {
	if e.isEmpty() {
		return false
	}
	return !(p.X() > e.Right() ||
		p.X() < e.Left() ||
		p.Y() > e.Top() ||
		p.Y() < e.Bottom())
}

func (e *Envelope) containsOther(other Envelope) bool {
	return e.coversOther(other)
}

func (e *Envelope) Covers(item interface{}) bool {
	return e.wrapBool(item, e.coversPoint, e.coversOther)
}

func (e *Envelope) coversOther(other Envelope) bool {
	if e.isEmpty() || other.isEmpty() {
		return false
	}
	return other.Left() >= e.Left() &&
		other.Right() <= e.Right() &&
		other.Top() <= e.Top() &&
		other.Bottom() >= e.Bottom()
}

func (e *Envelope) coversPoint(other Point) bool {
	if e.isEmpty() {
		return false
	}
	return other.X() >= e.Left() &&
		other.X() <= e.Right() &&
		other.X() <= e.Top() &&
		other.X() >= e.Bottom()
}

func (e *Envelope) Expand(item interface{}) {
	e.wrapVoid(item, e.expandForPoint, e.expandForOther)
}

func (e *Envelope) expandForPoint(p Point) {
	if e.isEmpty() {
		e[0] = p.X()
		e[1] = p.Y()
		e[2] = p.X()
		e[3] = p.Y()
	} else {
		if p.X() < e.Left() {
			e[0] = p.X()
		}
		if p.X() > e.Right() {
			e[2] = p.X()
		}
		if p.Y() > e.Top() {
			e[1] = p.Y()
		}
		if p.Y() < e.Bottom() {
			e[3] = p.Y()
		}
	}

}

func (e *Envelope) expandForOther(other Envelope) {
	if other.isEmpty() {
		return
	}
	if e.isEmpty() {
		e[0] = other.Left()
		e[1] = other.Top()
		e[2] = other.Right()
		e[3] = other.Bottom()
	} else {
		if other.Left() < e.Left() {
			e[0] = other.Left()
		}
		if other.Right() > e.Right() {
			e[2] = other.Right()
		}
		if other.Top() > e.Top() {
			e[1] = other.Top()
		}
		if other.Bottom() < e.Bottom() {
			e[3] = other.Bottom()
		}
	}

}
