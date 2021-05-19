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
	InOrderTraversal
)

// Iterator is a stateful iterator for traversing a linked list.
// Please note, that iteration over a list is not a thread-safe operation, and if parallel write operations are
// being performed on the list, the traversal can provide stale and outdated data.
type Iterator struct {
	currentNode    *Node
	traversalStack *lists.LinkedList
	traversalType  traversalType
	visitedNodes   map[*Node]bool
}

// endIterator is used for signifying the end of an iteration or traversal
var endIterator = &Iterator{}

// Next returns the iterator to the next node in the binary tree based on the traversal technique.
func (it *Iterator) Next() *Iterator {
	switch it.traversalType {
	case BreadthFirstTraversal:
		return bfsNext(it)
	case InOrderTraversal:
		return inorderNext(it)
	default:
		panic("invalid traversal type received")
	}
}

// Node returns the current node in the iteration
func (it *Iterator) Node() *Node {
	return it.currentNode
}

// BreadthFirstTraverse returns an iterator pointing to the root of the binary tree. The iterator is
// initialised in such a way that subsequent iterators (by calling Next()) returns tree nodes following the
// breadth-first traversal algorithm. Refer https://en.wikipedia.org/wiki/Breadth-first_search to know more.
func (t *Tree) BreadthFirstTraverse() *Iterator {
	if t.Root == nil {
		return endIterator
	}
	return &Iterator{currentNode: t.Root, traversalStack: lists.New(Node{}), traversalType: BreadthFirstTraversal}
}

// InOrderTraverse returns an iterator pointing to the root of the binary tree & is initialised in such a way that
// subsequent iterators (by calling Next()) returns tree nodes following In-order Traversal technique.
// For more info refer https://www.tutorialspoint.com/data_structures_algorithms/tree_traversal.htm
func (t *Tree) InOrderTraverse() *Iterator {
	if t.Root == nil {
		return endIterator
	}
	node := findLeftMost(t.Root)
	return &Iterator{currentNode: node, traversalType: InOrderTraversal, visitedNodes: map[*Node]bool{node: true}}
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
		currItr.traversalStack.PushBack(currItr.currentNode.Left)
	}
	if currItr.currentNode.Right != nil {
		currItr.traversalStack.PushBack(currItr.currentNode.Right)
	}

	if currItr.traversalStack.Empty() {
		return endIterator
	}
	currItr.currentNode = currItr.traversalStack.PopFront().(*Node)
	return currItr
}

// inorderNext is an auxiliary function that implements the in-order traversal algorithm.
func inorderNext(currItr *Iterator) *Iterator {
	currNode := currItr.currentNode
	rightMost := findRightMost(currNode)
	if rightMost != currNode {
		leftMost := findLeftMost(rightMost)
		currItr.currentNode = leftMost
	} else if currNode.parent != nil {
		_, found := currItr.visitedNodes[currNode.parent]
		if found {
			currItr.currentNode = currNode.parent.parent
		} else {
			currItr.currentNode = currNode.parent
		}
	}
	if currItr.currentNode == nil {
		return endIterator
	}
	currItr.visitedNodes[currItr.currentNode] = true
	return currItr
}

// findLeftMost returns the left-most child of the node, i.e., the node at the deepest left.
func findLeftMost(node *Node) *Node {
	var leftMost *Node
	if node.Left != nil {
		leftMost = findLeftMost(node.Left)
	} else {
		leftMost = node
	}
	return leftMost
}

// findRightMost returns the right-most child of the node, i.e., the node at the deepest right.
func findRightMost(node *Node) *Node {
	var rightMost *Node
	if node.Right != nil {
		rightMost = findRightMost(node.Right)
	} else {
		rightMost = node
	}
	return rightMost
}
