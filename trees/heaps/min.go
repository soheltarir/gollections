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
