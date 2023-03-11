package radixtree

import (
	"strings"
)

type Tree struct {
	label string
	edges []*Tree
}

func NewRadixTree() *Tree {
	return &Tree{label: "", edges: []*Tree{}}
}

func (r *Tree) isLeaf() bool {
	return len(r.edges) == 0
}

func (r *Tree) Lookup(x string) bool {
	node, offset, _ := r.query(x)
	return (node != nil && node.isLeaf() && offset == len(x))
}

func (r *Tree) Insert(x string) {
	node, offset, suffix := r.query(x)

	if offset == 0 {
		// add to node
		newLabel := x[offset:]
		newChild1 := &Tree{label: newLabel}
		node.edges = append(node.edges, newChild1)
	} else {
		// split the edge
		newParent := node.label[:len(suffix)]
		oldEdgeLabel := node.label[len(newParent):]
		newEdgeLabel := x[offset:]
		node.label = node.label[:len(suffix)]
		node.edges = []*Tree{{label: oldEdgeLabel, edges: node.edges}, {label: newEdgeLabel}}
	}
}

func (r *Tree) query(x string) (node *Tree, offset int, suffix string) {
	node = r
	length := len(x)
	for node != nil && !node.isLeaf() && offset < len(x) {
		for i := length; i > len(x[:offset]); i-- {
			suffix = x[offset:i]
			for _, edge := range node.edges {
				if strings.HasPrefix(edge.label, suffix) {
					node = edge
					offset += len(suffix)
					// when a edge need to be split
					if len(edge.label) > len(suffix) {
						return
					}
					continue
				}
			}
		}
	}

	return
}
