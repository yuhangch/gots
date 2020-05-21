package geom

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
