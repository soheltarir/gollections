package bst

import (
	"fmt"
	"github.com/soheltarir/gollections/core"
	"github.com/soheltarir/gollections/queue"
)

type Tree struct {
	root   *Node
	height int
	// internal attributes for type assertions
	expectedType string
	checker      core.TypeChecker
	comparator   core.Comparator
}

// Insert adds a new node at the leaf. Panics if type assertions fail
// - Time Complexity: O(log(n))
// - Space Complexity: O(log(n))
func (t *Tree) Insert(value interface{}) {
	// Do type assertions
	if t.checker != nil && !t.checker(value) {
		panic(fmt.Sprintf("invalid type provided, expected %s", t.expectedType))
	}
	newNode := &Node{Value: value}
	t.root, t.height = insertToTree(t.root, newNode, t.comparator, 0)
}

func (t *Tree) InsertMany(values ...interface{}) {
	for _, value := range values {
		t.Insert(value)
	}
}

// Height returns the height of the tree
// - Time Complexity: O(1)
// - Space Complexity: O(1)
func (t *Tree) Height() int {
	return t.height
}

// BreadthFirstSearch traverses the tree across breadth. For more refer: https://en.wikipedia.org/wiki/Breadth-first_search
//
//- Time Complexity: O(n)
//- Space Complexity: O(n)
func (t *Tree) BreadthFirstSearch() []interface{} {
	var nodes []interface{}

	// queue to store visited nodes
	q := queue.New()

	currentNode := t.root
	q.Enqueue(currentNode)

	for !q.Empty() {
		currentNode = q.Dequeue().(*Node)
		nodes = append(nodes, currentNode.Value)
		if currentNode.Left != nil {
			q.Enqueue(currentNode.Left)
		}
		if currentNode.Right != nil {
			q.Enqueue(currentNode.Right)
		}
	}
	return nodes
}

// NewInt returns a binary search tree with nodes containing int data
func NewInt() *Tree {
	return &Tree{
		comparator:   core.IntCompare,
		checker:      core.IntChecker,
		expectedType: "int",
	}
}

// NewString returns a binary search tree with nodes containing string data
func NewString() *Tree {
	return &Tree{
		expectedType: "string",
		checker:      core.StringChecker,
		comparator:   core.StringCompare,
	}
}
