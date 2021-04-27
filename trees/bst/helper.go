package bst

import (
	"github.com/soheltarir/gollections/core"
	"github.com/soheltarir/gollections/trees/binary_trees"
)

// insertToTree inserts a new node at the leaf and returns the updated height of the tree.
// This is a recursive helper function
func insertToTree(
	root *binary_trees.Node,
	newNode *binary_trees.Node,
	comparator core.Comparator,
	currHeight int,
) (*binary_trees.Node, int) {
	// Handle base case for recursion
	if root == nil {
		if currHeight == 0 {
			currHeight = 1
		}
		return newNode, currHeight
	}
	switch comparator(root.Value, newNode.Value) {
	case core.Greater:
		root.Left, currHeight = insertToTree(root.Left, newNode, comparator, currHeight)
		currHeight++
	case core.Lesser:
		root.Right, currHeight = insertToTree(root.Right, newNode, comparator, currHeight)
		currHeight++
	}
	return root, currHeight
}
