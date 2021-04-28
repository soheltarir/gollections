package heaps

import (
	"container/heap"
	"github.com/soheltarir/gollections/containers"
)

type _heap struct {
	heap.Interface
	data     []containers.Container
	size     int
	datatype containers.Container
}

func (h _heap) Len() int { return h.size }

func (h *_heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *_heap) Push(x interface{}) {
	h.data = append(h.data, h.datatype.Validate(x))
	h.size++
}

func (h *_heap) Pop() interface{} {
	popped := h.data[h.size-1]
	h.data = h.data[0 : h.size-1]
	h.size--
	return popped
}

// Insert allows both single & multiple elements to be added to the heap.
// Time complexity for adding a single element is O(log(n)).
// Time complexity for adding multiple elements is O(n)
func (h *_heap) Insert(values ...interface{}) {
	if len(values) == 1 {
		heap.Push(h, values[0])
	} else {
		var correctedValues []containers.Container
		var total int
		for _, value := range values {
			correctedValues = append(correctedValues, h.datatype.Validate(value))
			total++
		}
		h.data = append(h.data, correctedValues...)
		h.size = h.size + total
		heap.Init(h)
	}
}

func (h *_heap) Extract() interface{} {
	return heap.Pop(h)
}
