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

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/queue"
	"sync"
)

// TreeOperations creates a signature of common binary tree operations
type TreeOperations interface {
	// Insert adds an element in the tree
	Insert(interface{})
	// InsertMany add multiple elements in the tree
	InsertMany(...interface{})
}

// Tree defines the structure of a binary tree
type Tree struct {
	TreeOperations
	TreeIterations
	Root     *Node
	Height   int
	datatype containers.Container
	mu       sync.RWMutex
}

// Insert adds a node in the tree
func (t *Tree) Insert(value interface{}) {
	newNode := &Node{data: t.datatype.Validate(value)}

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Root == nil {
		t.Root = newNode
		t.Height++
		return
	}
	// Do a level order traversal until we find an empty place
	currNode := t.Root
	currHeight := 0

	q := queue.New(new(Node), currNode)
	for !q.Empty() {
		currNode = q.Dequeue().(*Node)
		currHeight++
		if currNode.Left == nil {
			currNode.Left = newNode
			newNode.parent = currNode
			if currHeight >= t.Height {
				t.Height++
			}
			break
		} else {
			q.Enqueue(currNode.Left)
		}
		if currNode.Right == nil {
			currNode.Right = newNode
			newNode.parent = currNode
			// We don't need to increment height here, as we are filling up the left leaf first.
			break
		} else {
			q.Enqueue(currNode.Right)
		}
	}
}

// InsertMany adds multiple nodes to the tree (order is preserved).
func (t *Tree) InsertMany(values ...interface{}) {
	for _, value := range values {
		t.TreeOperations.Insert(value)
	}
}

// NewInt instantiates a binary tree which can only accept integer as data.
func NewInt() *Tree {
	return New(containers.IntContainer(0))
}

// New instantiates a binary tree
func New(datatype containers.Container) *Tree {
	tree := &Tree{Root: nil, Height: 0, datatype: datatype, mu: sync.RWMutex{}}
	// The below handling is required to achieve method overriding.
	// Refer: https://stackoverflow.com/questions/38123911/golang-method-override
	tree.TreeOperations = interface{}(tree).(TreeOperations)
	tree.TreeIterations = interface{}(tree).(TreeIterations)
	return tree
}
