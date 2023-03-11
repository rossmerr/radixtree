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
	traverseNode := r
	elementsFound := 0
	for traverseNode != nil && !traverseNode.isLeaf() && elementsFound < len(x) {
		var nextEdge *Tree
		// Get the next edge to explore based on the elements not yet found in x
		//Edge nextEdge := select edge from traverseNode.edges where edge.label is a prefix of x.suffix(elementsFound)
		for _, edge := range traverseNode.edges {
			prefix := x[elementsFound:]
			if strings.HasPrefix(edge.label, prefix) {
				nextEdge = edge
			}
		}

		if nextEdge != nil {
			traverseNode = nextEdge

			elementsFound += len(nextEdge.label)
		} else {
			traverseNode = nil
		}

	}

	return (traverseNode != nil && traverseNode.isLeaf() && elementsFound == len(x))
}

func (r *Tree) Insert(x string) {
	traverseNode := r
	elementsFound := 0
	suffix := ""
	length := len(x)
	for traverseNode != nil && !traverseNode.isLeaf() && elementsFound < len(x) {
		var nextEdge *Tree
		for i := length; i > len(x[:elementsFound]); i-- {
			suffix = x[elementsFound:i]
			for _, edge := range traverseNode.edges {
				if strings.HasPrefix(edge.label, suffix) {
					// when a edge need to be split
					if len(edge.label) > len(suffix) {
						traverseNode = edge
						elementsFound += len(suffix)
						i = 0
						break
					}
					nextEdge = edge
					i = 0
					break
				}
			}
		}

		if nextEdge != nil {
			traverseNode = nextEdge
			elementsFound += len(suffix)
		} else {
			break
		}
	}

	if elementsFound == 0 {
		newLabel := x[elementsFound:]
		newChild1 := &Tree{label: newLabel}
		traverseNode.edges = append(traverseNode.edges, newChild1)
	} else {
		// split the edge
		newLabel := traverseNode.label[:len(suffix)]
		oldLabel1 := traverseNode.label[len(newLabel):]
		oldLabel2 := x[elementsFound:]

		newChild1 := &Tree{label: oldLabel1, edges: traverseNode.edges}
		newChild2 := &Tree{label: oldLabel2}

		traverseNode.label = newLabel
		traverseNode.edges = []*Tree{newChild1, newChild2}
	}
}
