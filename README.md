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
Below is an example `IntContainer` which implements the `Container` interface.
```go
func (c IntContainer) Key() interface{} {
	return int(c)
}

func (c IntContainer) Less(value Container) bool {
	y := value.(IntContainer)
	return int(c) < int(y)
}

func buildErrorMsg(expected interface{}, actual interface{}) string {
    return fmt.Sprintf("invalid type provided; expected: %s, received: %s",
        reflect.TypeOf(expected), reflect.TypeOf(actual))
}

func (IntContainer) Validate(x interface{}) Container {
	converted, ok := x.(int)
	if !ok {
		panic(buildErrorMsg(IntContainer(0), x))
	}
	return IntContainer(converted)
}
```

## Data Structures

---
Below are the data structures that are exposed by this package

- List
    - Linked Lists
    
- Maps
    - Counter: A map for counting hashable items. Sometimes called a bag or multiset. Elements are stored as map keys, and their counters are stored as map values
    
- Queue
- Stack
- Trees
    - Binary Trees
    - Binary Search Trees (BST)
    - Heaps