/**
MIT License

Copyright (c) 2021 Sohel Tarir

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package bst

import (
	"github.com/soheltarir/gollections/trees/binarytrees"
)

// insertToTree inserts a new node at the leaf and returns the updated height of the tree.
// This is a recursive helper function
func insertToTree(
	root *binarytrees.Node,
	newNode *binarytrees.Node,
	currHeight int,
) (*binarytrees.Node, int) {
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
