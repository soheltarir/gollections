package binary_trees

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/queue"
)

// TreeOperations creates a signature of common binary tree operations
type TreeOperations interface {
	InsertMany(...interface{})
	Insert(interface{})
}

// Tree defines the structure of a binary tree
type Tree struct {
	TreeOperations
	Root     *Node
	Height   int
	datatype containers.Container
}

// Insert adds a node in the tree
func (t *Tree) Insert(value interface{}) {

	newNode := &Node{Value: t.datatype.Validate(value)}
	if t.Root == nil {
		t.Root = newNode
		t.Height++
		return
	}
	// Do a level order traversal until we find an empty place
	currNode := t.Root
	currHeight := 0
	q := queue.New(currNode)
	for !q.Empty() {
		currNode = q.Dequeue().(*Node)
		currHeight++
		if currNode.Left == nil {
			currNode.Left = newNode
			if currHeight >= t.Height {
				t.Height++
			}
			break
		} else {
			q.Enqueue(currNode.Left)
		}
		if currNode.Right == nil {
			currNode.Right = newNode
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

func NewInt() *Tree {
	return New(containers.IntContainer(0))
}

// New instantiates a fresh binary tree
func New(datatype containers.Container) *Tree {
	tree := &Tree{Root: nil, Height: 0, datatype: datatype}
	// The below handling is required to achieve method overriding.
	// Refer: https://stackoverflow.com/questions/38123911/golang-method-override
	tree.TreeOperations = interface{}(tree).(TreeOperations)
	return tree
}