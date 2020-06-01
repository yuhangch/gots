package strtree

import "math"

const (
	DefaultNodeCapacity int = 10
)

type STRTree struct {
}

func (str *STRTree) verticalSlices(childBoundales Bounds, sliceCount int) []Bounds {
	cur := 0
	capacity := int(math.Ceil(float64(len(childBoundales)) / float64(sliceCount)))
	slices := make([]Bounds, sliceCount)
	for i := 0; i < sliceCount; i++ {
		for len(slices[i]) < capacity && cur < len(childBoundales) {
			slices[i] = append(slices[i], childBoundales[cur])
			cur++
		}
	}
	return slices

}
