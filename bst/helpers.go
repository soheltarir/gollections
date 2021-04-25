package bst

import (
	"github.com/soheltarir/gollections/stack"
	"github.com/soheltarir/gollections/utils"
)

func insertNode(root *Node, value utils.CollectionType) *Node {
	newNode := &Node{value: value}
	if root == nil {
		return newNode
	}
	switch root.value.Compare(value) {
	case -1:
		root.right = insertNode(root.right, value)
	case 1:
		root.left = insertNode(root.left, value)
	}
	return root
}

func breadthFirstSearch(root *Node) []utils.CollectionType {
	s := stack.New(root)
	var result []utils.CollectionType

	for !s.Empty() {
		root = s.Pop().(*Node)
		result = append(result, root.value)
		if root.left != nil {
			s.Push(root.left)
		}
		if root.right != nil {
			s.Push(root.right)
		}
	}
	return result
}
