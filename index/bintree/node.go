package bintree

import "fmt"

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
	if !(node == nil) {
		expand.expandToContain(node.interval)

	}
	larger := NewNodeByInterval(expand)
	if !(node == nil) {
		larger.insert(node)
	}
	return larger
}

func (nb *NodeBase) add(item interface{}) {
	nb.Items = append(nb.Items, item)
}
func (nb *NodeBase) addAll(items []interface{}) *[]interface{} {
	items = append(items, nb.Items...)
	for _, v := range nb.children {
		if v != nil {
			v.Items = append(v.Items, items...)
		}
	}
	return &items
}

func (nb *NodeBase) addAllOverlapingItems(interval Interval, result *[]interface{}) *[]interface{} {

	if !interval.isEmpty() && !nb.isSearchMatch(interval) {
		return result
	}

	*result = append(*result, nb.Items...)
	//fmt.Println(nb.Items,interval)

	for _, v := range nb.children {
		if v != nil {
			v.addAllOverlapingItems(interval, result)
		}
	}
	return result
}

func (nb *NodeBase) allChildren() (*Node, *Node) {
	return nb.children[0], nb.children[1]
}

func (nb *NodeBase) hasChildren() bool {
	for _, v := range nb.children {
		if v != nil {
			return true
		}
	}
	return false
}

func (nb *NodeBase) hasItems() bool {
	return len(nb.Items) > 0
}

func (nb *NodeBase) remove(itemInterval Interval, item interface{}) bool {
	if !nb.isSearchMatch(itemInterval) {
		return false
	}
	found := false
	for _, v := range nb.children {
		if v != nil {
			found = v.remove(itemInterval, item)
			if found {
				if v.isPrunable() {
					v = nil
				}
			}
		}

	}
	if found {
		return found
	}
	for k, v := range nb.Items {
		if v == item {
			nb.Items = append(nb.Items[0:k], nb.Items[k+1:len(nb.Items)-1]...)
			return true

		}
	}
	return false
}

func (nb *NodeBase) isPrunable() bool {
	return !(nb.hasChildren() || nb.hasItems())
}

func (nb *NodeBase) depth() int {
	childMaxDepth := 0
	for i := 0; i < 2; i++ {
		if nb.children[i] != nil {
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
		if nb.children[i] != nil {
			childSize += nb.children[i].size()
		}
	}

	return childSize + len(nb.Items)
}

func (nb *NodeBase) nodeSize() int {
	childSize := 0
	for i := 0; i < 2; i++ {
		if nb.children[i] != nil {
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
	root := &Root{

		origin: 0,
	}
	root.NodeBase.isSearchMatch = root.isSearchMatch
	return root
}

func (r *Root) insert(interval Interval, item interface{}) {
	index := childIndex(interval, r.origin)
	if index == -1 {
		r.add(item)
		return
	}
	child := r.children[index]
	if child == nil || !child.interval.contains(interval) {
		larger := createExpanded(child, interval)
		r.children[index] = larger
	}
	r.insertContained(r.children[index], interval, item)

}

//
func (r *Root) insertContained(tree *Node, itemIterval Interval, item interface{}) {
	if !tree.interval.contains(itemIterval) {
		fmt.Println("insert content")
		return
	}
	zero := itemIterval.width() == 0
	node := &Node{}
	if zero {
		node = tree.find(itemIterval)
	} else {
		node = tree.node(itemIterval)
	}
	node.add(item)

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
	node := &Node{
		interval: interval,
		level:    level,
		centre:   interval.centre(),
	}
	node.NodeBase.isSearchMatch = node.isSearchMatch
	return node
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
	if !(n.children[childIndex] == nil) {
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
	if node != nil {
		expand.expandToContain(node.interval)
	}
	larger := NewNodeByInterval(expand)
	if node != nil {
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
	if n.children[index] == nil {
		n.children[index] = n.createChild(index)
	}
	return n.children[index]
}
