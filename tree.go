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
	node, _, offset, _ := r.query(x)
	return (node != nil && node.isLeaf() && offset == len(x))
}

func (r *Tree) Insert(x string) {
	node, _, offset, suffix := r.query(x)

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

func (r *Tree) Delete(x string) bool {
	node, parent, _, _ := r.query(x)
	for i, edge := range parent.edges {
		if edge == node {
			parent.edges = append(parent.edges[:i], parent.edges[i+1:]...)
			return true
		}
	}

	return false
}

func (r *Tree) HasPrefix(x string) []string {
	node, _, _, _ := r.query(x)

	return hasPrefix(x, node)
}

func hasPrefix(prefix string, node *Tree) []string {

	results := []string{}
	if len(node.edges) > 0 {
		for _, edge := range node.edges {
			results = append(results, hasPrefix(prefix+edge.label, edge)...)
		}
	} else {
		results = append(results, prefix)
	}

	return results
}

func (r *Tree) query(x string) (node, parent *Tree, offset int, suffix string) {
	node = r
	length := len(x)
	for node != nil && !node.isLeaf() && offset < len(x) {
		parent = node
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
