package geom

import "math"

type Envelope [4]float64

func (e Envelope) Left() float64 {
	return e[0]
}

func (e Envelope) Top() float64 {
	return e[1]
}

func (e Envelope) Right() float64 {
	return e[2]
}

func (e Envelope) Bottom() float64 {
	return e[3]
}

func NewEmptyEnvelope() Envelope {
	return Envelope{0, -1, -1, 0}
}

func NewEnvelope(coordinates [4]float64) Envelope {
	var env Envelope = coordinates
	if env.Height() < 0 {
		env[1], env[3] = env[3], env[1]
	}
	if env.Width() < 0 {
		env[0], env[2] = env[2], env[0]
	}
	return coordinates
}

func (e *Envelope) isEmpty() bool {
	if e.Width() < 0 || e.Height() < 0 {
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
	c := NewCoordinate([]float64{x, y})
	return &c

}
