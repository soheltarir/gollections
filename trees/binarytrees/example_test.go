package binarytrees_test

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/trees/binarytrees"
)

// User implements containers.Container interface
type User struct {
	ID    string
	Name  string
	Email string
}

func (u User) Key() interface{} {
	return u.ID
}

func (u User) Less(x containers.Container) bool {
	convertedX := x.(User)
	return u.Name < convertedX.Name
}

func (u User) Validate(x interface{}) containers.Container {
	converted, ok := x.(User)
	if !ok {
		panic("type conversion failed")
	}
	return converted
}

func ExampleTree() {

	user1 := User{ID: "1", Name: "Steve Rogers", Email: "steve@avengers.com"}
	user2 := User{ID: "2", Name: "Tony Stark", Email: "tony@avengers.com"}
	user3 := User{ID: "3", Name: "Natasha Romanoff", Email: "natasha@avengers.com"}

	// Create a new binary tree specifying the data-type for the nodes. You can provide an empty or existing struct
	// object to define the container's data, i.e., in the below example you could of initialised the tree using any of
	// the following:
	// - binarytrees.New(User{})
	// - binarytrees.New(new(User))
	// - binarytrees.New(user1)
	tree := binarytrees.New(User{})

	// Insert multiple nodes in the binary tree
	tree.InsertMany(user1, user2, user3)

	// Print the height of the tree
	fmt.Println(tree.Height)

	// Create a binary tree with integer data
	intTree := binarytrees.NewInt()
	intTree.InsertMany(1, 2, 3, 4, 5)
	fmt.Println(intTree.Height)

	// Output:
	// 2
	// 3
}
