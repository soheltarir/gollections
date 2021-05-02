package containers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/** Test Struct container */
type TestStruct struct {
	ID string
}

func (t TestStruct) Key() interface{} {
	return t.ID
}

func (t TestStruct) Less(item Container) bool {
	y := item.(TestStruct)
	return t.ID < y.ID
}

func (t TestStruct) Validate(x interface{}) Container {
	converted, ok := x.(TestStruct)
	if !ok {
		panic("Invalid type")
	}
	return converted
}

func TestIntContainer_Key(t *testing.T) {
	a := IntContainer(10)
	assert.Equal(t, 10, a.Key())
}

func TestIntContainer_Less(t *testing.T) {
	a := IntContainer(10)
	b := IntContainer(20)
	assert.True(t, a.Less(b))
}

func TestIntContainer_Validate(t *testing.T) {
	empty := IntContainer(0)
	assert.Equal(t, IntContainer(5), empty.Validate(5))
	assert.Panics(t, func() { empty.Validate("test") })
}

func TestStringContainer_Key(t *testing.T) {
	a := StringContainer("hi")
	assert.Equal(t, "hi", a.Key())
}

func TestStringContainer_Less(t *testing.T) {
	a := StringContainer("app")
	b := StringContainer("apple")
	assert.True(t, a.Less(b))
}

func TestStringContainer_Validate(t *testing.T) {
	empty := StringContainer("")
	assert.Equal(t, StringContainer("hi"), empty.Validate("hi"))
	assert.Panics(t, func() { empty.Validate(5) })
}

func TestCleanBasicType(t *testing.T) {
	a := StringContainer("a")
	assert.Equal(t, "a", CleanBasicType(a))
	one := IntContainer(1)
	assert.Equal(t, 1, CleanBasicType(one))
	item := TestStruct{ID: "1"}
	assert.Equal(t, item, CleanBasicType(item))
}
