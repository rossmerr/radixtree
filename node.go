package radixtree

type node struct {
	label string
	child []*node
}

func newNode() *node {
	return &node{
		child: []*node{},
	}
}

func newNodeWithLabel(label string) *node {
	return &node{
		label: label,
		child: []*node{},
	}
}
