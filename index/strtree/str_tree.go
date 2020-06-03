package strtree

const (
	DefaultNodeCapacity int = 10
)

type STRTree struct {
	nodeCapacity int
}

func NewSTRTree(capacity int) *STRTree {
	if capacity < 1 {
		panic("Node capacity must be greater than 1")
	}
	return &STRTree{
		nodeCapacity: capacity,
	}
}

//func (str *STRTree) verticalSlices(childBoundales Bounds, sliceCount int) []Bounds {
//	cur := 0
//	capacity := int(math.Ceil(float64(len(childBoundales)) / float64(sliceCount)))
//	slices := make([]Bounds, sliceCount)
//	for i := 0; i < sliceCount; i++ {
//		for len(slices[i]) < capacity && cur < len(childBoundales) {
//			slices[i] = append(slices[i], childBoundales[cur])
//			cur++
//		}
//	}
//	return slices
//}
//
//
//func (str *STRTree) createParentBoundables(childBoundales Bounds, newLevel int) {
//	if len(childBoundales)<0 {
//		panic("空")
//	}
//	leafCount := int(math.Ceil(math.Sqrt(float64(len(childBoundales))/float64(str.nodeCapacity))))
//	var parentBoundales []*Node
//	parentBoundales = append(parentBoundales, NewNode(newLevel))
//	childBoundales.SortY()
//	slices := str.verticalSlices(childBoundales,leafCount)
//	return str.createParentBoundablesFromVerticalSlices(slices,newLevel)
//
//}
//
//func (str *STRTree) createParentBoundablesFromVerticalSlices(vs []Bounds,newLevel int) {
//	if len(vs) < 1 {
//		panic("")
//	}
//	var parentBoundales []*Node
//	for _,v:=range vs {
//		parentBoundales = append(parentBoundales,str.createParentBoundablesFromVerticalSlice(v,newLevel))
//	}
//
//}
//
//func (str *STRTree) createParentBoundablesFromVerticalSlice(b Bounds,newLevel int) {
//
//}
//
//func (str *STRTree) parentBoundables(childBounds Bounds,newLevel int) Nodes {
//	if len(childBounds)<1 {
//		panic("")
//	}
//	var parentBoundales Nodes
//	parentBoundales = append(parentBoundales, NewNode(newLevel))
//	childBounds.SortY()
//	for _,v:=range childBounds{
//		if len(parentBoundales.lastNode().childBoundales) == str.nodeCapacity {
//			parentBoundales = append(parentBoundales, NewNode(newLevel))
//		}
//		parentBoundales.lastNode().appendChildBounds(v)
//	}
//	return parentBoundales
//
//}
