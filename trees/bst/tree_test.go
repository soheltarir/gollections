package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInt(t *testing.T) {
	tree := NewInt()
	// Positive Test
	tree.Insert(1)
	// Negative test, should panic
	assert.Panics(t, func() {
		tree.Insert("one")
	}, "Should panic with 'one'")
}

func TestNewString(t *testing.T) {
	tree := NewString()
	// Positive Test
	tree.Insert("one")
	// Negative test, should panic
	assert.Panics(t, func() { tree.Insert(1) }, "Should panic with 1")
}

func TestTree_Insert(t *testing.T) {
	tree := NewInt()
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(11)
	assert.Equal(t, 2, tree.Height())
}

func TestTree_InsertMany(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(10, 1, 11)
	assert.Equal(t, 2, tree.Height())
}

func TestTree_BreadthFirstSearch(t *testing.T) {
	tree := NewInt()
	tree.InsertMany(10, 1, 11, 3, 4, 12)
	// Need to convert the result to int slice first
	res := tree.BreadthFirstSearch()
	var actual []int
	for _, val := range res {
		actual = append(actual, val.(int))
	}
	assert.Equal(t, []int{10, 1, 11, 3, 12, 4}, actual)
}
