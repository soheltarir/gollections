package containers

import (
	"fmt"
	"reflect"
)

type Container interface {
	Key() interface{}
	Less(container Container) bool
	Validate(interface{}) Container
}

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
