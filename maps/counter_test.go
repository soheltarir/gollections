package maps

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStringCounter(t *testing.T) {
	counter := NewStringCounter()
	counter.Add("a")
	counter.Add("a")
	assert.Equal(t, map[interface{}]int{"a": 2}, counter.counterMap)
	assert.Panics(t, func() { counter.Add(1) })
}

func TestNewIntCounter(t *testing.T) {
	counter := NewIntCounter()
	counter.Add(1)
	counter.Add(1)
	assert.Equal(t, map[interface{}]int{1: 2}, counter.counterMap)
	assert.Panics(t, func() { counter.Add("a") })
}

func TestCounter_Subtract(t *testing.T) {
	counter := NewStringCounter()
	counter.Add("a")
	counter.Add("a")
	counter.Subtract("a")
	assert.Equal(t, map[interface{}]int{"a": 1}, counter.counterMap)
	counter.Subtract("c")
	assert.Equal(t, -1, counter.Get("c"))
}

func TestCounter_MostCommon(t *testing.T) {
	counter := NewStringCounter()
	counter.Add("a")
	counter.Add("a")
	counter.Add("b")
	expected := map[containers.Container]int{containers.StringContainer("a"): 2}
	assert.Equal(t, expected, counter.MostCommon(1))
	counter.Add("b")
	expected = map[containers.Container]int{containers.StringContainer("a"): 2, containers.StringContainer("b"): 2}
	assert.Equal(t, expected, counter.MostCommon(2))
}

func TestCounter_Iterator(t *testing.T) {
	counter := NewStringCounter()
	arr := []interface{}{"a", "a", "b", "c", "d", "a", "c"}
	counter.AddMany(arr...)
	counter.Iterator(func(key interface{}, value int) {
		assert.Equal(t, counter.counterMap[key], value)
	})
}
