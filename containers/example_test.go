package containers_test

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
)

// User implements containers.Container interface
type User struct {
	ID    string
	Name  string
	Email string
}

/**
Any struct/type should implement the following methods to consider it a `Container`.
*/

// Key signifies a unique identifier for the struct/type. This function should return the datatype which
// implement equality (=) operator, i.e., the datatype of the Key should be one of Go's basic types (https://tour.golang.org/basics/11)
func (u User) Key() interface{} {
	return u.ID
}

// Less whether the corresponding User is object is less than the target User
func (u User) Less(x containers.Container) bool {
	convertedX := x.(User)
	return u.Name < convertedX.Name
}

// Validate converts an interface type to User container
// Panics if an invalid interface is provided
func (u User) Validate(x interface{}) containers.Container {
	converted, ok := x.(User)
	if !ok {
		// Handle for pointer to struct objects
		converted2, ok2 := x.(*User)
		if !ok2 {
			panic("type conversion failed")
		}
		return converted2
	}
	return converted
}

func Example() {
	user1 := User{ID: "1", Name: "A", Email: "a@example.com"}
	user2 := User{ID: "2", Name: "B", Email: "b@example.com"}

	fmt.Println("User 1 Key: ", user1.Key())
	fmt.Println("User 2 Key: ", user2.Key())
	fmt.Println("Is User 1 Less than User 2? ", user1.Less(user2))

	// Output:
	// User 1 Key:  1
	// User 2 Key:  2
	// Is User 1 Less than User 2?  true
}
