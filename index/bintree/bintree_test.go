package bintree

import (
	"fmt"
	"math"
	"testing"
)

const (
	numItems = 20000.0
	minExt   = -1000.0
	maxExt   = 1000.0
)

var (
	tree      = NewBinTree()
	intervals []Interval
)

func initArgs() (float64, float64, float64, float64) {
	gridSize := math.Floor(math.Sqrt(numItems)) + 1
	extent := maxExt - minExt
	delta := extent / gridSize
	cellSize := 2 * delta

	return gridSize, extent, delta, cellSize
}

func initTree() {
	gs, _, del, cs := initArgs()
	for i := 0; i < int(gs); i++ {
		x := minExt + del*float64(i)
		interval := NewInterval(x, x+cs)
		tree.insert(interval, interval)
		//fmt.Println(tree.size())
		intervals = append(intervals, interval)
		//fmt.Println(len(intervals))
	}
}

func queryGrid(n int, cellSize float64, t *testing.T) {
	gridSize := math.Floor(math.Sqrt(float64(n))) + 1
	extent := maxExt - minExt
	delta := extent / gridSize
	for i := 0; i < int(gridSize); i++ {
		x := minExt + delta*float64(i)
		query(NewInterval(x, x+cellSize), t)
	}
}

func query(interval Interval, t *testing.T) {
	candidates := tree.queryByInterval(interval)
	results := overlapping(candidates, interval)

	validResults := queryIntervals(interval)

	//fmt.Println(len(results), "results")
	//fmt.Println(len(validResults), "valid results")
	if len(results) != len(validResults) {
		t.Error("failed")
	}

}

func overlapping(items *[]interface{}, search Interval) []interface{} {
	var result []interface{}
	for _, v := range *items {
		if v == nil {
			continue
		}
		i := v.(Interval)
		if i.overlaps(search) {
			result = append(result, i)
		}
	}
	return result
}

func queryIntervals(search Interval) (results []Interval) {
	for _, v := range intervals {
		if v.overlaps(search) {
			results = append(results, v)
		}
	}
	return
}

func TestBT(t *testing.T) {
	_, _, _, cs := initArgs()
	initTree()
	fmt.Printf("depth = %d  size = %d\n", tree.depth(), tree.size())
	queryGrid(100, cs, t)
}
