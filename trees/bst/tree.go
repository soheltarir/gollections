package bst

import (
	"fmt"
	"github.com/soheltarir/gollections/core"
	"github.com/soheltarir/gollections/queue"
	"github.com/soheltarir/gollections/trees/binary_trees"
)

type Tree struct {
	binary_trees.Tree
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
	newNode := &binary_trees.Node{Value: value}
	t.Root, t.Height = insertToTree(t.Root, newNode, t.comparator, 0)
}

// BreadthFirstSearch traverses the tree across breadth. For more refer: https://en.wikipedia.org/wiki/Breadth-first_search
//
//- Time Complexity: O(n)
//- Space Complexity: O(n)
func (t *Tree) BreadthFirstSearch() []interface{} {
	var nodes []interface{}

	// queue to store visited nodes
	q := queue.New()

	currentNode := t.Root
	q.Enqueue(currentNode)

	for !q.Empty() {
		currentNode = q.Dequeue().(*binary_trees.Node)
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
	tree := &Tree{
		comparator:   core.IntCompare,
		checker:      core.IntChecker,
		expectedType: "int",
	}
	// The below handling is required to achieve method overriding.
	// Refer: https://stackoverflow.com/questions/38123911/golang-method-override
	tree.TreeOperations = interface{}(tree).(binary_trees.TreeOperations)
	return tree
}

// NewString returns a binary search tree with nodes containing string data
func NewString() *Tree {
	tree := &Tree{
		expectedType: "string",
		checker:      core.StringChecker,
		comparator:   core.StringCompare,
	}
	// The below handling is required to achieve method overriding.
	// Refer: https://stackoverflow.com/questions/38123911/golang-method-override
	tree.TreeOperations = interface{}(tree).(binary_trees.TreeOperations)
	return tree
}
