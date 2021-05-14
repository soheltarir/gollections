package main

import (
	"fmt"
	"github.com/soheltarir/gollections/trees/binarytrees"
)

func main() {
	tree := binarytrees.NewInt()
	tree.InsertMany(1, 2, 3, 4, 11, 5, 6, 7)
	tree.Insert(10)
	for it := tree.BeginBreadthFirstSearch(); it != tree.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}
