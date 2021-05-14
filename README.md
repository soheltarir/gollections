# Gollections

---

![Tests](https://github.com/soheltarir/gollections/actions/workflows/unittest.yml/badge.svg)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

A Go library which provides common programming data structures and functions like C++ STL library, and specialized container datatypes like Python's collection module.

## Containers

---
Container is the base interface that most of the data structures use. 
Any struct/type should implement the methods to be considered as a `Container`.
```go
type Container interface {
	// Key signifies a unique identifier for the struct/type. 
	// This function should return the datatype which implement 
	// equality (=) operator, i.e., the datatype of the Key should be one of 
	// Go's basic types (https://tour.golang.org/basics/11)
	Key() interface{}
	// Less reports whether the implementing struct/type is less than the target 
	// Container or not.
	Less(container Container) bool
	// Validate converts an interface type to the corresponding Container. 
	// Should panic if an invalid interface is provided.
	Validate(interface{}) Container
}
```
Below is an example `User` struct which implements the `Container` interface.
```go
package main

import (
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
```

## Data Structures

### Basic Example

Using the `User` container used as an example in the above section, the below implementation demonstrates the usage of
data-structures exposed by the package.

```go
package main

import (
  "fmt"
  "github.com/soheltarir/gollections/queue"
)

func main() {
  // Initialise a Queue
  q := queue.New(User{})
  // Add element to queue
  user1 := &User{ID: "1", Name: "John Wick", Email: "john.wick@example.com"}
  q.Enqueue(user1)

  // Add another user to queue
  user2 := &User{ID: "2", Name: "Hanzo Hashashi", Email: "scorpion@example.com"}
  q.Enqueue(user2)

  // Retrieve the next user in queue
  fmt.Println(q.Dequeue())
}
```
Below is the list of data-structures exposed by this package. All the below mentioned data-structures provide
thread-safe operations.

- [Lists](https://pkg.go.dev/github.com/soheltarir/gollections/lists)
- [Queue](https://pkg.go.dev/github.com/soheltarir/gollections/queue)
- [Stack](https://pkg.go.dev/github.com/soheltarir/gollections/stack)
- [Maps](https://pkg.go.dev/github.com/soheltarir/gollections/maps)
    
    - [Counter](https://pkg.go.dev/github.com/soheltarir/gollections/maps/counter)
