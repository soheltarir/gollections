package containers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
