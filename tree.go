package radixtree

import (
	"math"
	"strings"
)

type Tree struct {
	root *node
}

func NewTree(label string) *Tree {
	return &Tree{
		root: newNodeWithLabel(""),
	}
}

func (s *Tree) Insert(label string) {
	s.insert(label, s.root)
}

func (s *Tree) insert(label string, currentNode *node) {
	iterator, newLabel, _, _ := s.enumerate(label, currentNode)
	inserted := false
	for iterator.HasNext() {
		child, _ := iterator.Next()
		inserted = true
		s.insert(newLabel, child)

	}

	if !inserted {
		currentNode.child = append(currentNode.child, newNodeWithLabel(newLabel))
	}

}

func (s *Tree) Lookup(label string) bool {
	return s.lookup(label, s.root)
}

func (s *Tree) lookup(label string, currentNode *node) bool {
	iterator, newLabel, matches, matched := s.enumerate(label, currentNode)

	for iterator.HasNext() {
		child, _ := iterator.Next()
		return s.lookup(newLabel, child)

	}
	if !matched && matches == len(currentNode.label) {
		return true
	}

	return false

}

func (s *Tree) Successor(label string) string {
	return s.successor(label, s.root, "")
}

func (s *Tree) successor(label string, currentNode *node, carry string) string {
	matches := matchingConsecutiveCharacters(label, currentNode)

	if (matches == 0) || (currentNode == s.root) ||
		((matches > 0) && (matches < len(label))) {

		newLabel := label[matches : len(label)-matches]
		for _, child := range currentNode.child {
			if strings.HasPrefix(child.label, string(newLabel[0])) {
				return s.successor(newLabel, child, carry+currentNode.label)
			}
		}
		return currentNode.label
	} else if matches < len(currentNode.label) {
		return carry + currentNode.label
	} else if matches == len(currentNode.label) {
		carry = carry + currentNode.label

		min := math.MaxInt
		index := -1
		for i := 0; i < len(currentNode.child); i++ {
			if len(currentNode.child[i].label) < min {
				min = len(currentNode.child[i].label)
				index = i
			}
		}

		if index > -1 {
			return carry + currentNode.child[index].label
		} else {
			return carry
		}

	}
	return ""

}

func (s *Tree) Predecessor(label string) string {
	return s.predecessor(label, s.root, "")
}

func (s *Tree) predecessor(label string, currentNode *node, carry string) string {
	iterator, newLabel, matches, matched := s.enumerate(label, currentNode)

	for iterator.HasNext() {
		child, _ := iterator.Next()
		return s.predecessor(newLabel, child, carry+currentNode.label)

	}
	if !matched && matches == len(currentNode.label) {
		return carry + currentNode.label
	}

	return ""
}

func (s *Tree) Delete(label string) {
	s.delete(label, s.root)
}

func (s *Tree) delete(label string, currentNode *node) {
	iterator, newLabel, _, _ := s.enumerate(label, currentNode)
	for iterator.HasNext() {
		child, index := iterator.Next()
		if newLabel == child.label {
			if len(child.child) == 0 {
				currentNode.child = append(currentNode.child[:index], currentNode.child[index+1:]...)
				return
			}
		}
	}
}

func (s *Tree) enumerate(label string, currentNode *node) (*treeIterator, string, int, bool) {
	return NewTreeIterator(s, currentNode, label)
}

type treeIterator struct {
	currentNode *node
	indexStart  int
	indexEnd    int
	label       string
	matched     bool
}

func NewTreeIterator(tree *Tree, currentNode *node, label string) (*treeIterator, string, int, bool) {

	matches := matchingConsecutiveCharacters(label, currentNode)

	newLabel := label[matches : len(label)-matches]

	matched := (matches == 0) || (currentNode == tree.root) ||
		((matches > 0) && (matches < len(label)) && (matches >= len(currentNode.label)))

	return &treeIterator{
		indexStart:  0,
		indexEnd:    len(currentNode.child),
		currentNode: currentNode,
		label:       newLabel,
	}, newLabel, matches, matched
}

func (s *treeIterator) HasNext() bool {
	return s.matched && s.indexStart < s.indexEnd
}

func (s *treeIterator) Next() (child *node, index int) {
	for index, child = range s.currentNode.child[s.indexStart:] {
		s.indexStart = index
		if strings.HasPrefix(child.label, string(s.label[0])) {
			return
		}
	}
	return
}

func matchingConsecutiveCharacters(label string, currentNode *node) int {
	matches := 0
	minLength := 0

	if len(currentNode.label) >= len(label) {
		minLength = len(label)
	} else if len(currentNode.label) < len(label) {
		minLength = len(currentNode.label)
	}

	if minLength > 0 {
		//go throught the two streams
		for i := 0; i < minLength; i++ {
			//if two characters at the same position have the same value we have one more match
			if label[i] == currentNode.label[i] {
				matches++
			} else {
				//if at any position the two strings have different characters break the cycle
				break
			}
		}
	}
	//and return the current number of matches
	return matches
}
