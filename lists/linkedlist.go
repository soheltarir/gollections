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

// Package lists exposes adaptors for doubly-linked list data-structure.
//
// Lists are sequence containers that allow constant time insert and erase operations anywhere within the sequence,
// and iteration in both directions.
// List containers are implemented as doubly-linked lists; Doubly linked lists can store each of the elements they
// contain in different and unrelated storage locations. The ordering is kept internally by the association to each
// element of a link to the element preceding it and a link to the element following it.
package lists

import (
	"fmt"
	"strings"
	"sync"

	"github.com/soheltarir/gollections/containers"
)

// Node represents an element in a Linked List
type Node struct {
	Value    containers.Container
	next     *Node
	previous *Node
}

// LinkedList is a sequence container that allow constant time insert and erase operations anywhere within the sequence,
// and iteration in both directions.
type LinkedList struct {
	head      *Node
	tail      *Node
	size      int64
	valueType containers.Container
	mu        sync.RWMutex
}

/** Element Access **/

// Front returns the value of the first element of the linked list
func (ll *LinkedList) Front() interface{} {
	head := ll.head
	if head == nil {
		return nil
	}
	// handle for nilNode
	return containers.CleanBasicType(head.Value)
}

// Back returns the value of the last element of the linked list
func (ll *LinkedList) Back() interface{} {
	tail := ll.tail
	if tail == nil {
		return nil
	}
	return containers.CleanBasicType(tail.Value)
}

/** Iterators */

// Begin returns an iterator pointing to the first element in the list container.
func (ll *LinkedList) Begin() *Iterator {
	ll.mu.RLock()
	defer ll.mu.RUnlock()

	if ll.size == 0 {
		return ll.End()
	}
	head := ll.head
	return &Iterator{currentNode: head, direction: forwardDirection, index: 0, limit: ll.size - 1}
}

// End Returns an iterator referring to the past-the-end element in the list container.
func (ll *LinkedList) End() *Iterator {
	return endFwdIterator
}

// RBegin returns a reverse iterator pointing to the last element in the container (i.e., its reverse beginning).
// Reverse iterators iterate backwards: increasing them moves them towards the beginning of the container.
func (ll *LinkedList) RBegin() *Iterator {
	if ll.size == 0 {
		return ll.REnd()
	}
	tail := ll.tail
	return &Iterator{currentNode: tail, direction: backwardDirection, index: ll.size - 1, limit: ll.size - 1}
}

// REnd returns a reverse iterator pointing to the theoretical element preceding the first element
// in the list container
func (ll *LinkedList) REnd() *Iterator {
	return endBackIterator
}

/** Modifiers */

// PushFront inserts a new element at the beginning of the list, right before its current first element.
// This effectively increases the container size by one.
// Panics if an invalid type is provided.
func (ll *LinkedList) PushFront(val interface{}) {
	element := ll.valueType.Validate(val)
	node := &Node{Value: element}

	ll.mu.Lock()
	defer ll.mu.Unlock()

	if ll.size == 0 {
		ll.head, ll.tail = node, node
	} else {
		tmpNode := ll.head
		tmpNode.previous = node
		node.next = tmpNode
		node.previous = nil
		ll.head = node
	}
	ll.size++
}

// PushBack adds a new element at the end of the list container, after its current last element.
// This effectively increases the container size by one.
// Panics if an invalid type is provided.
func (ll *LinkedList) PushBack(val interface{}) {
	element := ll.valueType.Validate(val)
	node := &Node{Value: element}

	ll.mu.Lock()
	defer ll.mu.Unlock()

	if ll.size == 0 {
		ll.head, ll.tail = node, node
	} else {
		tail := ll.tail
		tail.next = node
		node.previous = tail
		node.next = nil
		ll.tail = node
	}
	ll.size++
}

// PopFront deletes the first element of the list and returns it's value
func (ll *LinkedList) PopFront() interface{} {
	if ll.size == 0 {
		return nil
	}

	ll.mu.Lock()
	defer ll.mu.Unlock()

	head := ll.head
	next := head.next
	if next != nil {
		next.previous = nil
	}
	ll.head = next
	ll.size--
	return containers.CleanBasicType(head.Value)
}

// PopBack deletes the last element of the list and returns it's value
func (ll *LinkedList) PopBack() interface{} {
	if ll.size == 0 {
		return nil
	}

	ll.mu.Lock()
	defer ll.mu.Unlock()

	tail := ll.tail
	previous := tail.previous
	previous.next = nil
	ll.tail = previous
	ll.size--
	return containers.CleanBasicType(tail.Value)
}

// Insert extends the list by inserting new elements before the element at the specified position.
// This effectively increases the list size by the amount of elements inserted.
// Note: This operation is not thread safe.
func (ll *LinkedList) Insert(it *Iterator, elements ...interface{}) {
	tempList := New(ll.valueType)

	for _, element := range elements {
		tempList.PushBack(element)
	}
	if ll.size == 0 {
		ll.head, ll.tail = tempList.head, tempList.tail
		ll.size = tempList.size
	} else {
		tmpHead, tmpTail := tempList.head, tempList.tail
		prev := it.currentNode.previous
		if prev != nil {
			prev.next = tmpHead
			tmpHead.previous = prev
		} else {
			// This is the case wherein the elements are being pushed before the head of the list
			ll.head = tmpHead
		}
		it.currentNode.previous = tmpTail
		tmpTail.next = it.currentNode
		ll.size += tempList.size
	}
}

// Clear Removes all elements from the list container (which are destroyed), and leaving the list with a size of 0.
func (ll *LinkedList) Clear() {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	ll.head, ll.tail, ll.size = nil, nil, 0
}

func (ll *LinkedList) eraseSingle(it *Iterator) {
	prev := it.currentNode.previous
	next := it.currentNode.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.previous = prev
	}
	// Edge Handling
	if it.currentNode == ll.head {
		ll.head = next
	}
	if it.currentNode == ll.tail {
		ll.tail = prev
	}
	ll.size--
}

// Erase removes from the list container either a single element or a range of elements ([first,last)).
// Note: The bounds are including the first iterator & excluding the last iterator
func (ll *LinkedList) Erase(iterators ...*Iterator) error {
	if len(iterators) > 2 || len(iterators) == 0 {
		return fmt.Errorf("please provide a single iterator or the iterator bounds (i.e., only two iterators)")
	}
	ll.mu.Lock()
	defer ll.mu.Unlock()

	if len(iterators) == 1 {
		it := iterators[0]
		ll.eraseSingle(it)
		return nil
	}

	first, last := iterators[0], iterators[1]

	for it := first; !it.IsEqual(last); it = it.Next() {
		ll.eraseSingle(it)
	}
	return nil
}

/** Capacity Functions **/

// Size returns the length of the linked list
func (ll *LinkedList) Size() int64 {
	return ll.size
}

// Empty returns whether the list container is empty (i.e. whether its size is 0).
func (ll *LinkedList) Empty() bool {
	return ll.size == 0
}

/** Display Functions **/

//Display returns a string representation of the linked list.
func (ll *LinkedList) Display() string {
	var b strings.Builder

	for it := ll.Begin(); it != ll.End(); it = it.Next() {
		if it.currentNode != ll.tail {
			_, _ = fmt.Fprintf(&b, "%v <-> ", it.currentNode.Value.Key())
		} else {
			_, _ = fmt.Fprintf(&b, "%v", it.currentNode.Value.Key())
		}
	}
	return b.String()
}

/** Constructors **/

// New constructs an empty container linked list, with no elements.
func New(valueType containers.Container) *LinkedList {
	return &LinkedList{valueType: valueType, mu: sync.RWMutex{}, size: 0}
}

// NewInt constructs an empty integer linked list, with no elements.
func NewInt() *LinkedList {
	return New(containers.IntContainer(0))
}
