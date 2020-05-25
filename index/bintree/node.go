package bintree

type NodeBase struct {
	Items         []interface{}
	children      [2]*Node
	isSearchMatch func(interval Interval) bool
}

func childIndex(itemInterval Interval, centre float64) int {
	index := -1
	if itemInterval.min() >= centre {
		index = 1
	}
	if itemInterval.max() <= centre {
		index = 0
	}
	return index
}

func createExpanded(node *Node, add Interval) *Node {
	expand := NewInterval(add.min(), add.max())
	if !node.isEmpty() {
		expand.expandToContain(node.interval)

	}
	larger := NewNodeByInterval(expand)
	if !node.isEmpty() {
		larger.insert(node)
	}
	return larger
}

func (nb *NodeBase) add(item interface{}) {
	nb.Items = append(nb.Items, item)
}
func (nb *NodeBase) addAll(items []interface{}) []interface{} {
	items = append(items, nb.Items...)
	for i := 0; i < 2; i++ {
		if !nb.children[i].isEmpty() {
			nb.children[i].Items = append(nb.children[i].Items, items...)
		}
	}
	return items
}

func (nb *NodeBase) addAllOverlapingItems(interval Interval, result []interface{}) {
	if !interval.isEmpty() && nb.isSearchMatch(interval) {
		return
	}

	result = append(result, nb.Items...)
	if !nb.children[0].isEmpty() {
		nb.children[0].addAllOverlapingItems(interval, result)
	}
	if !nb.children[1].isEmpty() {
		nb.children[1].addAllOverlapingItems(interval, result)
	}

}

func (nb *NodeBase) remove(itemInterval Interval, item interface{}) bool {
	if !nb.isSearchMatch(itemInterval) {
		return false
	}
	found := false
	for i := 1; i < 2; i++ {
		if !nb.children[i].isEmpty() {
			found = nb.children[i].remove(itemInterval, item)
		}
		if found {
			if nb.children[i].isPrunable() {

			}
		}
	}
}

func (nb *NodeBase) isPrunable() bool {

}

func (nb *NodeBase) depth() int {
	childMaxDepth := 0
	for i := 0; i < 2; i++ {
		if !nb.children[i].isEmpty() {
			childDpt := nb.children[i].depth()
			if childDpt > childMaxDepth {
				childMaxDepth = childDpt
			}
		}
	}
	return childMaxDepth + 1
}

func (nb *NodeBase) size() int {
	childSize := 0
	for i := 0; i < 2; i++ {
		if !nb.children[i].isEmpty() {
			childSize += nb.children[i].size()
		}
	}

	return childSize + len(nb.Items)
}

func (nb *NodeBase) nodeSize() int {
	childSize := 0
	for i := 0; i < 2; i++ {
		if !nb.children[i].isEmpty() {
			childSize += nb.children[i].nodeSize()
		}
	}
	return childSize + 1
}

type Root struct {
	NodeBase
	origin float64
}

func NewRoot() *Root {
	return &Root{
		origin: 0,
	}
}

func (r *Root) isEmpty() bool {
	return true
}

func (r *Root) insert(interval Interval, item interface{}) {
	index := childIndex(interval, r.origin)
	if index == -1 {
		r.add(item)
		return
	}
	child := r.children[index]
	if child.isEmpty() || !child.interval.contains(interval) {
		larger := createExpanded(child, interval)
		r.children[index] = larger
	}
	r.insertContained(r.children[index], interval, item)

}

//
func (r *Root) insertContained(tree *Node, itemIterval Interval, item interface{}) {
	if tree.interval.contains(itemIterval) {
		panic("")
	}
	zero := itemIterval.width() < 0.0001
	node := &Node{}
	if zero {
		node = tree.find(itemIterval)
	} else {
		node = tree.node(itemIterval)
	}
	node.add(itemIterval)

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

func NewNodeByInterval(interval Interval) *Node {
	key := NewKey(interval)
	return NewNode(key.interval, key.level)

}

func (n *Node) isEmpty() bool {
	return n.level == 0 && n.centre == 0
}

func (n *Node) node(search Interval) *Node {
	childIndex := childIndex(search, n.centre)
	if childIndex != -1 {
		node := n.child(childIndex)
		return node.node(search)
	} else {
		return n
	}
}

func (n *Node) find(search Interval) *Node {
	childIndex := childIndex(search, n.centre)
	if childIndex == -1 {
		return n
	}
	if !n.children[childIndex].isEmpty() {
		node := n.child(childIndex)
		return node.find(search)
	}
	return n
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

func (n *Node) createExpanded(node *Node, interval Interval) *Node {
	expand := NewInterval(interval.min(), interval.max())
	if node.isEmpty() {
		expand.expandToContain(node.interval)
	}
	larger := NewNodeByInterval(expand)
	if node.isEmpty() {
		larger.insert(node)
	}
	return larger
}

func (n *Node) insert(node *Node) {
	if !(n.interval.isEmpty() || n.interval.contains(node.interval)) {
		panic("")
	}
	index := childIndex(node.interval, n.centre)
	if node.level == n.level-1 {
		n.children[index] = node
	} else {
		child := n.createChild(index)
		child.insert(node)
		n.children[index] = child
	}
}

func (n *Node) child(index int) *Node {
	if n.children[index].isEmpty() {
		n.children[index] = n.createChild(index)
	}
	return n.children[index]
}
