package containers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	assert.Equal(t, 10, ToInt(IntContainer(10)))
	assert.Panics(t, func() { ToInt(StringContainer("hi")) })
}

func TestToString(t *testing.T) {
	assert.Equal(t, "hi", ToString(StringContainer("hi")))
	assert.Panics(t, func() { ToString(IntContainer(10)) })
}

func TestToIntSlice(t *testing.T) {
	arr := []Container{IntContainer(10), IntContainer(2), IntContainer(3), IntContainer(5)}
	expected := []int{10, 2, 3, 5}
	assert.Equal(t, expected, ToIntSlice(arr))
}

func TestToStringSlice(t *testing.T) {
	arr := []Container{StringContainer("hello"), StringContainer("world")}
	expected := []string{"hello", "world"}
	assert.Equal(t, expected, ToStringSlice(arr))
}
