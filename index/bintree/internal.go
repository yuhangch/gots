package bintree

type Interval [2]float64

func (i *Interval) min() float64 {
	return i[0]
}

func (i *Interval) max() float64 {
	return i[1]
}

func (i *Interval) width() float64 {
	return i.max() - i.min()
}

func (i *Interval) centre() float64 {
	return (i.max() + i.min()) / 2
}

func (i *Interval) update(min, max float64) {
	if max < min {
		min, max = max, min
	}
	i[0], i[1] = min, max
}

func (i *Interval) overlaps(t Interval) bool {
	return !(i.min() > t.max() || i.max() < t.min())
}

func (i *Interval) contains(iter Interval) bool {
	return i.min() <= iter.min() && i.max() >= iter.max()
}

func (i *Interval) expandToContain(item Interval) {
	if item.max() > i.max() {
		i[1] = item.max()
	}
	if item.min() < i.min() {
		i[0] = item.min()
	}
}

func NewInterval(min, max float64) Interval {
	if max < min {
		min, max = max, min
	}
	return Interval{min, max}
}
