package strtree

import "github.com/yuhangch/gots/geom"

type Node struct {
	boundaries []*geom.Envelope
	boundary   *geom.Envelope
	level      int
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
	return len(n.boundaries)
}

func (n *Node) append(b *geom.Envelope) {
	n.boundaries = append(n.boundaries, b)
}

type Nodes []*Node

func (n *Nodes) lastNode() *Node {
	return (*n)[len(*n)-1]
}

type Leaf struct {
	Node
	Items geom.GeometryCollection
}
