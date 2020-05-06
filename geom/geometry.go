package geom

type Geometry interface {
	Type() string
	Envelope() Envelope
}
