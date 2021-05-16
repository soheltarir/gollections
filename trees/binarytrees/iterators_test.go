package binarytrees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_InOrderTraversal(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(10, 9, 6, 5, 11, 20)
	expected := []interface{}{5, 9, 11, 10, 20, 6}
	assert.Equal(t, expected, tree.InOrderTraversal())
}

func TestTree_PreOrderTraversal(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(10, 9, 6, 5, 11, 20)
	expected := []interface{}{10, 9, 5, 11, 6, 20}
	assert.Equal(t, expected, tree.PreOrderTraversal())
}

func TestTree_PostOrderTraversal(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(10, 9, 6, 5, 11, 20)
	expected := []interface{}{5, 11, 9, 20, 6, 10}
	assert.Equal(t, expected, tree.PostOrderTraversal())
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
