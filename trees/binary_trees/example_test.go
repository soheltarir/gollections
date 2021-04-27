package binary_trees_test

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/trees/binary_trees"
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

func Example() {
	user1 := User{ID: "1", Name: "Steve Rogers", Email: "steve@avengers.com"}
	user2 := User{ID: "2", Name: "Tony Stark", Email: "tony@avengers.com"}
	user3 := User{ID: "3", Name: "Natasha Romanoff", Email: "natasha@avengers.com"}

	tree := binary_trees.New(User{})
	tree.InsertMany(user1, user2, user3)
	fmt.Println(tree.Height)

	// Output:
	// 2
}
