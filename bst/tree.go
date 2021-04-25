package bst

import "github.com/soheltarir/gollections/utils"

type IntTree struct {
	root	*Node
}

func NewInt() *IntTree {
	return &IntTree{}
}

func (t *IntTree) Insert(value int) {
	val := utils.IntType(value)
	t.root = insertNode(t.root, val)
}

func (t *IntTree) BreadthTraverse() []int {
	result := breadthFirstSearch(t.root)
	var intResult []int
	for _, val := range result {
		intResult = append(intResult, int(val.(utils.IntType)))
	}
	return intResult
}
