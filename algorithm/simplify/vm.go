package simplify

import (
	"container/list"
	"github.com/yuhangch/gots/geom"
	"math"
)

type AreaChain list.List

type AreaNode struct {
	Index int
	Area  float64
}

type VmCfg struct {
	Line      *geom.LineString
	Tolerance float64
	Chain     *list.List
	Removed   *map[int]bool
}

func newAreaNode(index int, area float64) AreaNode {
	return AreaNode{
		Index: index,
		Area:  area}
}

func buildAreaChain(line geom.LineString) *list.List {
	chain := list.New()
	for k := 1; k < len(line)-1; k++ {
		A, B, C := (line)[k-1], (line)[k], (line)[k+1]
		area := geom.Triangle{A, B, C}.Area()
		chain.PushBack(newAreaNode(k, area))
	}
	chain.PushFront(newAreaNode(0, math.Inf(1)))
	chain.PushBack(newAreaNode(len(line)-1, math.Inf(1)))
	return chain

}

func recomputeArea(cfg *VmCfg, node *list.Element) {
	nodeValue := node.Value.(AreaNode)
	if nodeValue.Area == math.Inf(1) {
		return
	}
	nodeValue.Area = triangle(cfg, node).Area()
	cfg.Chain.InsertAfter(nodeValue, node)
	cfg.Chain.Remove(node)

}

func triangle(cfg *VmCfg, node *list.Element) geom.Triangle {
	index := node.Value.(AreaNode).Index
	prevs := node.Prev().Value.(AreaNode).Index
	next := node.Next().Value.(AreaNode).Index

	return geom.Triangle{(*cfg.Line)[prevs], (*cfg.Line)[index], (*cfg.Line)[next]}

}

func simplifyNode(cfg *VmCfg) bool {
	min := cfg.Tolerance
	var element *list.Element
	for e := (*cfg.Chain).Front(); e != nil; e = e.Next() {
		if area := e.Value.(AreaNode).Area; area < min {
			min = area
			element = e
		}
	}
	if element != nil {
		p, l := element.Prev(), element.Next()
		(*cfg.Chain).Remove(element)
		recomputeArea(cfg, p)
		recomputeArea(cfg, l)
		(*cfg.Removed)[element.Value.(AreaNode).Index] = true
		return true

	}
	return false
}

func VM(lineString geom.LineString, tolerance float64) geom.LineString {
	var simplified geom.LineString
	remove := make(map[int]bool, len(lineString))
	cfg := &VmCfg{
		Line:      &lineString,
		Tolerance: tolerance,
		Chain:     buildAreaChain(lineString),
		Removed:   &remove,
	}

	for {
		if !simplifyNode(cfg) {
			break
		}
	}
	for k := range lineString {
		if _, ok := (*cfg.Removed)[k]; !ok {
			simplified = append(simplified, lineString[k])
		}
	}
	return simplified
}
