package strtree

import "github.com/yuhangch/gots/geom"

type Node struct {
	childBoundales []interface{}
	bounds         interface{}
	level          int
}

func (n *Node) computeBounds() geom.Envelope {

}

func (n *Node) Size() int {
	return len(n.childBoundales)
}
