package bst

import (
	"github.com/soheltarir/gollections/trees/binary_trees"
)

// insertToTree inserts a new node at the leaf and returns the updated height of the tree.
// This is a recursive helper function
func insertToTree(
	root *binary_trees.Node,
	newNode *binary_trees.Node,
	currHeight int,
) (*binary_trees.Node, int) {
	// Handle base case for recursion
	if root == nil {
		if currHeight == 0 {
			currHeight = 1
		}
		return newNode, currHeight
	}
	if root.Value.Less(newNode.Value) {
		root.Right, currHeight = insertToTree(root.Right, newNode, currHeight)
		currHeight++
	} else {
		root.Left, currHeight = insertToTree(root.Left, newNode, currHeight)
		currHeight++
	}
	return root, currHeight
}
