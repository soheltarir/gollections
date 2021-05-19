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

// Package binarytrees exposes the Binary Tree container, refer https://en.wikipedia.org/wiki/Binary_tree to know more
// about the container.
// Binary Trees are data-structures in which each node has at most two children, which are referred to as the
// left child and the right child.
package binarytrees

import "github.com/soheltarir/gollections/containers"

// Node is the basic building block of a binary tree. It contains the following:
// 1. Data
// 2. Pointer to left child
// 3. Pointer to right child
type Node struct {
	data   containers.Container
	Left   *Node
	Right  *Node
	parent *Node
}

// Data returns the data stored in the tree node
func (n Node) Data() interface{} {
	return containers.CleanBasicType(n.data)
}

func (n Node) Key() interface{} {
	return n.data.Key()
}

// Less compares the node with another binary tree node, and compares them.
func (n Node) Less(x containers.Container) bool {
	y := x.(Node)
	return n.data.Less(y.data)

}

// Validate checks whether the interface provided is either a Node or *Node.
// Panics if the check fails.
func (Node) Validate(x interface{}) containers.Container {
	converted, ok := x.(Node)
	if !ok {
		return x.(*Node)
	}
	return converted
}

func NewTreeNode(data containers.Container) *Node {
	return &Node{data: data}
}
