package containers

import (
	"fmt"
	"reflect"
)

// Container is base interface that most of the data structures use.
// Any struct/type should implement the methods to be considered as a `Container`.
// Refer IntContainer for implementation reference
type Container interface {
	// Key signifies a unique identifier for the struct/type. This function should return the datatype which
	// implement equality (=) operator, i.e., the datatype of the Key should be one of Go's basic types (https://tour.golang.org/basics/11)
	Key() interface{}
	// Less reports whether the implementing struct/type is less than the target Container or not.
	Less(container Container) bool
	// Validate converts an interface type to the corresponding Container. Should panic if an invalid interface is provided.
	Validate(interface{}) Container
}

// IntContainer is a Container encapsulation over basic int type
type IntContainer int

func (c IntContainer) Key() interface{} {
	return int(c)
}

func (c IntContainer) Less(value Container) bool {
	y := value.(IntContainer)
	return int(c) < int(y)
}

func (IntContainer) Validate(x interface{}) Container {
	converted, ok := x.(int)
	if !ok {
		panic(buildErrorMsg(IntContainer(0), x))
	}
	return IntContainer(converted)
}

// StringContainer is a Container encapsulation over the basic string type
type StringContainer string

func (c StringContainer) Key() interface{} {
	return string(c)
}

func (c StringContainer) Less(value Container) bool {
	x := string(c)
	y := string(value.(StringContainer))
	return x < y
}

func (StringContainer) Validate(x interface{}) Container {
	converted, ok := x.(string)
	if !ok {
		panic(buildErrorMsg(StringContainer(""), x))
	}
	return StringContainer(converted)
}

// buildErrorMsg builds an error string to inform the difference between the expected & actual value/s.
func buildErrorMsg(expected interface{}, actual interface{}) string {
	return fmt.Sprintf("invalid type provided; expected: %s, received: %s",
		reflect.TypeOf(expected), reflect.TypeOf(actual))
}
