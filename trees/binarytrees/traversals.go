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

package binarytrees

import "github.com/soheltarir/gollections/lists"

/**
Recursive Auxiliary functions
*/

type traversalType uint

const (
	breadthFirstTraversal = iota
)

type Iterator struct {
	currentNode   *Node
	visited       *lists.LinkedList
	traversalType traversalType
}

var endIterator = &Iterator{}

func (it *Iterator) Next() *Iterator {
	if it.currentNode == nil {
		return endIterator
	}
	if it.currentNode.Left != nil {
		it.visited.PushBack(it.currentNode.Left)
	}
	if it.currentNode.Right != nil {
		it.visited.PushBack(it.currentNode.Right)
	}

	if it.visited.Empty() {
		return endIterator
	}
	it.currentNode = it.visited.PopFront().(*Node)
	return it
}

func (it *Iterator) Value() interface{} {
	return it.currentNode.Value
}

func (t *Tree) BeginBreadthFirstSearch() *Iterator {
	return &Iterator{currentNode: t.Root, visited: lists.New(Node{})}
}

func (t *Tree) End() *Iterator {
	return endIterator
}

// inorderTraversalAuxiliary is a recursive function to traverse the tree InOrder depth first
func inorderTraversalAuxiliary(node *Node, result []interface{}) []interface{} {
	if node.Left != nil {
		result = inorderTraversalAuxiliary(node.Left, result)
	}
	result = append(result, node.Value.Key())
	if node.Right != nil {
		result = inorderTraversalAuxiliary(node.Right, result)
	}
	return result
}

// preorderTraversalAuxiliary is a recursive function to traverse the tree PreOrder depth first
func preorderTraversalAuxiliary(node *Node, result []interface{}) []interface{} {
	result = append(result, node.Value.Key())
	if node.Left != nil {
		result = preorderTraversalAuxiliary(node.Left, result)
	}
	if node.Right != nil {
		result = preorderTraversalAuxiliary(node.Right, result)
	}
	return result
}

// postorderTraversalAuxiliary is a recursive function to traverse the tree PostOrder depth first
func postorderTraversalAuxiliary(node *Node, result []interface{}) []interface{} {
	if node.Left != nil {
		result = postorderTraversalAuxiliary(node.Left, result)
	}
	if node.Right != nil {
		result = postorderTraversalAuxiliary(node.Right, result)
	}
	result = append(result, node.Value.Key())
	return result
}

/*****************************************************************************************************/

func (t *Tree) InOrderTraversal() []interface{} {
	var result []interface{}
	return inorderTraversalAuxiliary(t.Root, result)
}

func (t *Tree) PreOrderTraversal() []interface{} {
	var result []interface{}
	return preorderTraversalAuxiliary(t.Root, result)
}

func (t *Tree) PostOrderTraversal() []interface{} {
	var result []interface{}
	return postorderTraversalAuxiliary(t.Root, result)
}
