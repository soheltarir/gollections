package heaps

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMinInt(t *testing.T) {
	heap := NewMinInt(1, 10, 3, 6)
	assert.IsType(t, containers.IntContainer(0), heap.data[0])
	assert.Equal(t, 1, containers.ToInt(heap.data[0]))
}

func TestNewMaxInt(t *testing.T) {
	heap := NewMaxInt(1, 10, 3, 6)
	assert.IsType(t, containers.IntContainer(0), heap.data[0])
	assert.Equal(t, 10, containers.ToInt(heap.data[0]))
}

func Test_heap_Insert(t *testing.T) {
	heap := NewMaxInt()
	heap.Insert(10)
	heap.Insert(1)
	heap.Insert(20)
	assert.Equal(t, 20, containers.ToInt(heap.data[0]))
}

func Test_heap_Extract(t *testing.T) {
	heap := NewMinInt(10, 20, 30, 5)
	assert.Equal(t, 5, containers.ToInt(heap.Extract()))
}
