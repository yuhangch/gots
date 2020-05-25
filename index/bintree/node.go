package bintree

type NodeBase struct {
	Items         []interface{}
	children      [2]*Node
	isSearchMatch func(Interval)
}

func subNodeIndex(itemInterval Interval, centre float64) int {
	index := -1
	if itemInterval.min() >= centre {
		index = 1
	}
	if itemInterval.max() <= centre {
		index = 0
	}
	return index
}

func (nb *NodeBase) add(item interface{}) {
	nb.Items = append(nb.Items, item)
}
func (nb *NodeBase) addAll(items []interface{}) []interface{} {
	items = append(items, nb.Items...)

	return items
}

type Root struct {
	NodeBase
}

func (r *Root) insertContained(tree Node, itemIterval Interval, item interface{}) {
	if tree.interval.contains(itemIterval) {
		panic("")
	}

}

func (r *Root) isSearchMatch(item Interval) bool {
	return true
}

type Node struct {
	NodeBase
	interval Interval
	centre   float64
	level    int
}

func NewNode(interval Interval, level int) *Node {
	return &Node{
		interval: interval,
		level:    level,
		centre:   interval.centre(),
	}
}

func (n *Node) isEmpty() bool {
	return n.level == 0 && n.centre == 0
}

func (n *Node) isSearchMatch(item Interval) bool {
	return item.overlaps(n.interval)
}

func (n *Node) createChild(index int) (child *Node) {
	min, max := 0.0, 0.0
	switch index {
	case 0:
		min = n.interval.min()
		max = n.centre
		break
	case 1:
		min = n.centre
		max = n.interval.max()
		break
	}
	return NewNode(NewInterval(min, max), n.level-1)
}

func (n *Node) child(index int) *Node {
	if n.children[index].isEmpty() {
		n.children[index] = n.createChild(index)
	}
	return n.children[index]
}
