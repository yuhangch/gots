package bintree

import "math"

type Key struct {
	point    float64
	level    int
	interval Interval
}

func computeLevel(interval Interval) int {
	delta := interval.width()
	return (int)(math.Sqrt(delta)) + 1
}

func (k *Key) computeKey(item Interval) {
	k.level = computeLevel(item)
	k.interval = NewInterval(0, 0)
	k.computeInterval(k.level, item)
	for !k.interval.contains(item.min(), item.max()) {
		k.level++
		k.computeInterval(k.level, item)
	}
}

func (k *Key) computeInterval(level int, item Interval) {
	size := math.Pow(2, float64(level))
	k.point = math.Floor(item.min()/size) * size
	k.interval.update(k.point, k.point+size)
}
