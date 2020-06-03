package strtree

import "github.com/yuhangch/gots/geom"

type Node struct {
	childBoundales []interface{}
	bounds         interface{}
	level          int
}

func NewNode(level int) *Node {
	return &Node{
		level: level,
	}
}

func (n *Node) computeBounds() *geom.Envelope {
	return geom.NewEmptyEnvelope()
}

func (n *Node) Size() int {
	return len(n.childBoundales)
}

func (n *Node) appendChildBounds(b interface{}) {
	n.childBoundales = append(n.childBoundales, b)
}

type Nodes []*Node

func (n *Nodes) lastNode() *Node {
	return (*n)[len(*n)-1]
}
