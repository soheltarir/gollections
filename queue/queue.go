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

// Package queue exposes adaptors for FIFO (first-in first-out) data-structure.
//
// Queues are a type of container adaptor, specifically designed to operate in a FIFO context (first-in first-out),
// where elements are inserted into one end of the container and extracted from the other.
//
// Elements are pushed into the "back" of the specific container and popped from its "front".
package queue

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/lists"
)

// A Queue is a linear structure which follows a particular order in which the operations are performed.
// The order is First In First Out (FIFO)
type Queue struct {
	data *lists.LinkedList
}

// Enqueue Inserts a new element at the end of the queue, after its current last element.
func (q *Queue) Enqueue(value interface{}) {
	q.data.PushBack(value)
}

// Dequeue Removes the next element in the queue, effectively reducing its size by one.
// The element removed is the "oldest" element in the queue whose value can be retrieved by calling method Front().
func (q *Queue) Dequeue() interface{} {
	return q.data.PopFront()
}

// Front Returns the next element in the queue.
func (q Queue) Front() interface{} {
	return q.data.Front()
}

// Back Returns the last element in the queue.
// This is the "newest" element in the queue (i.e. the last element pushed into the queue).
func (q Queue) Back() interface{} {
	return q.data.Back()
}

// Size returns the total size of the queue
func (q *Queue) Size() int64 {
	return q.data.Size()
}

// Empty returns true if the queue has no items
func (q *Queue) Empty() bool {
	return q.data.Empty()
}

// Clear empties the queue
func (q *Queue) Clear() {
	q.data.Clear()
}

// New instantiates a new queue with the items provided (order is preserved)
func New(valueType containers.Container, values ...interface{}) *Queue {
	// Initialise a linked list
	list := lists.New(valueType)
	list.Insert(list.Begin(), values...)
	return &Queue{data: list}
}

// NewInt instantiates a new queue which can contain integer elements
func NewInt(values ...interface{}) *Queue {
	return New(containers.IntContainer(0), values...)
}

// NewString instantiates a new queue which can contain string elements
func NewString(values ...interface{}) *Queue {
	return New(containers.StringContainer(""), values...)
}
