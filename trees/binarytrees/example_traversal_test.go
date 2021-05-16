package binarytrees_test

import (
	"fmt"
	"github.com/soheltarir/gollections/trees/binarytrees"
)

func ExampleTree_BreadthFirstTraverse() {
	// Create a new binary tree containing integer data
	tree := binarytrees.NewInt()
	tree.InsertMany(1, 2, 4, 5, 6, 3)

	// Iterate through the tree using breadth-first search technique, and print the node data.
	for it := tree.BreadthFirstTraverse(); it != tree.End(); it = it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// 1
	// 2
	// 4
	// 5
	// 6
	// 3
}
