/**
MIT License

Copyright (c) 2021 Sohel Tarir

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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

// CleanBasicType checks whether the container provided is one of Go's basic types and converts them accordingly
// to avoid manual conversion from containers (implemented by the package).
func CleanBasicType(container Container) interface{} {
	switch container.(type) {
	case IntContainer:
		return container.Key()
	case StringContainer:
		return container.Key()
	default:
		return container
	}
}
