package counter

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStringCounter(t *testing.T) {
	counter := NewStringCounter()
	counter.Add("a")
	counter.Add("a")
	assert.Equal(t, 2, counter.Get("a"))
	assert.Panics(t, func() { counter.Add(1) })
}

func TestNewIntCounter(t *testing.T) {
	counter := NewIntCounter()
	counter.Add(1)
	counter.Add(1)
	assert.Equal(t, 2, counter.Get(1))
	assert.Panics(t, func() { counter.Add("a") })
}

func TestCounter_Subtract(t *testing.T) {
	counter := NewStringCounter()
	counter.Add("a")
	counter.Add("a")
	counter.Subtract("a")
	assert.Equal(t, 1, counter.Get("a"))
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
	counter.Range(func(key interface{}, value int) {
		assert.Equal(t, counter.Get(key), value)
	})
}

func TestCounter_Delete(t *testing.T) {
	counter := NewStringCounter("a", "a", "b")
	counter.Delete("a")
	assert.Equal(t, 0, counter.Get("a"))
	assert.Equal(t, 1, counter.Size())
}
