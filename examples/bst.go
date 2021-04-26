package main

import (
	"fmt"
	"github.com/soheltarir/gollections/trees/bst"
)

func main() {
	btree := bst.NewInt()
	btree.Insert(5)
	btree.Insert(10)
	btree.Insert(1)
	btree.Insert(11)
	btree.Insert(9)
	btree.Insert(4)
	fmt.Println(btree.Height())
	fmt.Println(btree.BreadthFirstSearch())
}
