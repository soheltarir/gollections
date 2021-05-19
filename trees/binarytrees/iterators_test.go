package binarytrees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_InOrderTraversal(t *testing.T) {
	tree := NewInt()
	var expected []interface{}
	var actual []interface{}

	// Test for empty tree
	for it := tree.InOrderTraverse(); it != tree.End(); it = it.Next() {
		actual = append(actual, it.Node().Data())
	}
	assert.Equal(t, expected, actual)

	actual = make([]interface{}, 0)

	tree.InsertMany(10, 9, 6, 5, 11, 20)
	expected = []interface{}{5, 9, 11, 10, 20, 6}

	for it := tree.InOrderTraverse(); it != tree.End(); it = it.Next() {
		actual = append(actual, it.Node().Data())
	}
	assert.Equal(t, expected, actual)
}

func TestIterator_Next(t *testing.T) {
	tree := NewInt()
	it := tree.BreadthFirstTraverse()
	assert.Equal(t, endIterator, it.Next())

	it = &Iterator{currentNode: nil, traversalType: 10}
	assert.Panics(t, func() {
		it.Next()
	})
}

func TestLeftMost(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(1, 2, 3, 4, 5, 6, 7)
	assert.Equal(t, 4, findLeftMost(tree.Root).Data())
}

func TestRightMost(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(1, 2, 3, 4, 5, 6, 7)
	assert.Equal(t, 7, findRightMost(tree.Root).Data())
}
