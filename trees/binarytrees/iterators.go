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

// TreeIterations lists methods to iterate through a binary tree.
type TreeIterations interface {
	// BreadthFirstTraverse returns an iterator pointing to the root of the binary tree. The iterator is
	// initialised in such a way that subsequent iterators (by calling Next()) returns tree nodes following the
	// breadth-first traversal algorithm. Refer https://en.wikipedia.org/wiki/Breadth-first_search to know more.
	BreadthFirstTraverse() *Iterator
}

// traversalType defines enums for tree traversal techniques
type traversalType uint

const (
	BreadthFirstTraversal traversalType = iota
)

// Iterator is a stateful iterator for traversing a linked list.
// Please note, that iteration over a list is not a thread-safe operation, and if parallel write operations are
// being performed on the list, the traversal can provide stale and outdated data.
type Iterator struct {
	currentNode   *Node
	visited       *lists.LinkedList
	traversalType traversalType
}

// endIterator is used for signifying the end of an iteration or traversal
var endIterator = &Iterator{}

// Next returns the iterator to the next node in the binary tree based on the traversal technique.
func (it *Iterator) Next() *Iterator {
	switch it.traversalType {
	case BreadthFirstTraversal:
		return bfsNext(it)
	default:
		panic("invalid traversal type received")
	}
}

// Value returns the current node's data
func (it *Iterator) Value() interface{} {
	return it.currentNode.Value
}

// BreadthFirstTraverse returns an iterator pointing to the root of the binary tree. The iterator is
// initialised in such a way that subsequent iterators (by calling Next()) returns tree nodes following the
// breadth-first traversal algorithm. Refer https://en.wikipedia.org/wiki/Breadth-first_search to know more.
func (t *Tree) BreadthFirstTraverse() *Iterator {
	return &Iterator{currentNode: t.Root, visited: lists.New(Node{})}
}

// End returns an iterator to the past-the-end node in the binary tree.
func (t *Tree) End() *Iterator {
	return endIterator
}

// bfsNext is an auxiliary function that implements the breadth-first search traversal algorithm.
func bfsNext(currItr *Iterator) *Iterator {
	if currItr.currentNode == nil {
		return endIterator
	}
	if currItr.currentNode.Left != nil {
		currItr.visited.PushBack(currItr.currentNode.Left)
	}
	if currItr.currentNode.Right != nil {
		currItr.visited.PushBack(currItr.currentNode.Right)
	}

	if currItr.visited.Empty() {
		return endIterator
	}
	currItr.currentNode = currItr.visited.PopFront().(*Node)
	return currItr
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
