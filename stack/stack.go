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

// Package stack exposes adaptors for LIFO (last-in first-out) data-structure.
//
// Stacks are a type of container adaptor, specifically designed to operate in a LIFO context (last-in first-out),
// where elements are inserted and extracted only from one end of the container.
package stack

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/lists"
)

// Stack is a type of container adaptor, specifically designed to operate in a LIFO context (last-in first-out),
// where elements are inserted and extracted only from one end of the container.
type Stack struct {
	data *lists.LinkedList
}

// Push Inserts a new element at the top of the stack, above its current top element.
func (s *Stack) Push(value interface{}) {
	s.data.PushBack(value)
}

// Pop Removes the element on top of the stack, effectively reducing its size by one.
// The element removed is the latest element inserted into the stack, whose value can be retrieved by calling
// method Stack::Top.
func (s *Stack) Pop() interface{} {
	return s.data.PopBack()
}

// Size Returns the number of elements in the stack.
func (s *Stack) Size() int64 {
	return s.data.Size()
}

// Empty Returns whether the stack is empty: i.e. whether its size is zero.
func (s *Stack) Empty() bool {
	return s.data.Empty()
}

// Top Returns a reference to the top element in the stack.
// Since stacks are last-in first-out containers, the top element is the last element inserted into the stack.
func (s *Stack) Top() interface{} {
	return s.data.Back()
}

// Clear deletes all the elements in the stack, effectively reducing its size to 0
func (s *Stack) Clear() {
	s.data.Clear()
}

// New instantiates a fresh stack with the values provided
func New(valueType containers.Container, values ...interface{}) *Stack {
	list := lists.New(valueType)
	list.Insert(list.Begin(), values...)
	return &Stack{data: list}
}

// NewInt constructs a stack containing only integer elements
func NewInt(values ...interface{}) *Stack {
	return New(containers.IntContainer(0), values...)
}

// NewString constructs a stack containing only string elements
func NewString(values ...interface{}) *Stack {
	return New(containers.StringContainer(""), values...)
}
