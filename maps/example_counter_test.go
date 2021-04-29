package maps_test

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/maps"
)

// User implements containers.Container interface
type User struct {
	ID   string
	Name string
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

func Example_counter() {
	user1 := User{ID: "1", Name: "Steve Rogers"}
	user2 := User{ID: "2", Name: "Tony Stark"}
	user3 := User{ID: "3", Name: "Natasha Romanoff"}

	// Initialise the Counter. Pass an empty struct to set the datatype of the counter
	// You can send elements in the NewCounter function itself.
	counter := maps.NewCounter(User{}, user1, user2, user3)
	// Increment the counter for the object
	counter.Add(user1)
	// Increment counters for multiple objects
	counter.AddMany(user2, user3)
	// Each subtract decreases the counter for the object. The count can be less than zero.
	counter.Subtract(user2)

	// Print the count of the elements
	fmt.Println(counter.Get(user1))
	fmt.Println(counter.Get(user2))

	// Completely removes the key from the counter
	counter.Delete(user2)
	fmt.Println(counter.Get(user2))

	// Prints the size of the Counter object
	fmt.Println(counter.Size())

	// Output:
	// 2
	// 1
	// 0
	// 2
}
