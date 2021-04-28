package heaps

import (
	"container/heap"
	"github.com/soheltarir/gollections/containers"
)

/** Min Heap */

type MinHeap struct {
	_heap
}

func (h MinHeap) Less(i, j int) bool {
	return h.data[i].Less(h.data[j])
}

func NewMin(datatype containers.Container, elements ...interface{}) *MinHeap {
	h := &MinHeap{_heap{datatype: datatype}}
	h.Interface = interface{}(h).(heap.Interface)
	if len(elements) > 0 {
		h.Insert(elements...)
	}
	return h
}

func NewMinInt(elements ...interface{}) *MinHeap {
	return NewMin(containers.IntContainer(0), elements...)
}

/** Max Heap */

type MaxHeap struct {
	_heap
}

func (h MaxHeap) Less(i, j int) bool {
	return h.data[j].Less(h.data[i])
}

// NLargest returns a list with n largest elements
func (h MaxHeap) NLargest(n int) []containers.Container {
	if n > h.size {
		n = h.size
	}
	return h.data[0:n]
}

func NewMax(datatype containers.Container, elements ...interface{}) *MaxHeap {
	h := &MaxHeap{_heap{datatype: datatype}}
	h.Interface = interface{}(h).(heap.Interface)
	if len(elements) > 0 {
		h.Insert(elements...)
	}
	return h
}

func NewMaxInt(elements ...interface{}) *MaxHeap {
	return NewMax(containers.IntContainer(0), elements...)
}
