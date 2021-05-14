package binary_trees

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tree := NewInt()
	assert.IsType(t, &Tree{}, tree)
}

func TestTree_Insert(t *testing.T) {
	tree := NewInt()
	tree.Insert(10)
	assert.Equal(t, 10, tree.Root.Value.Key())
}

func TestTree_Height(t *testing.T) {
	tree := NewInt()
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(9)
	assert.Equal(t, 3, tree.Height)
}

func TestTree_InsertMany(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(10, 1, 11, 15, 6)
	assert.Equal(t, 3, tree.Height)
}

func ExampleTree_Insert() {
	// Initialise a new integer binary tree
	tree := NewInt()
	// Insert a value in the tree
	tree.Insert(10)
}

func TestNode_Key(t *testing.T) {
	node := Node{Value: containers.IntContainer(1)}
	assert.Equal(t, 1, node.Key())
}

func TestNode_Less(t *testing.T) {
	node1 := Node{Value: containers.IntContainer(1)}
	node2 := Node{Value: containers.IntContainer(2)}
	assert.True(t, node1.Less(node2))
}

func TestNode_Validate(t *testing.T) {
	tester := Node{Value: containers.StringContainer("")}
	nodeNormal := Node{Value: containers.StringContainer("hello")}
	nodePointer := &Node{Value: containers.StringContainer("world")}
	assert.Equal(t, nodeNormal, tester.Validate(nodeNormal))
	assert.Equal(t, nodePointer, tester.Validate(nodePointer))
	assert.Panics(t, func() {
		tester.Validate("a")
	})
}
