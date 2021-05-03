package lists

import (
	"github.com/soheltarir/gollections/containers"
	"sync/atomic"
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
	head      atomic.Value //*Node
	tail      atomic.Value //*Node
	size      int64
	valueType containers.Container
}

/** Iterators */

// Begin returns an iterator pointing to the first element in the list container.
func (ll *LinkedList) Begin() *Iterator {
	head := ll.head.Load().(*Node)
	return &Iterator{currentNode: head, direction: forwardDirection, index: 0}
}

// End Returns an iterator referring to the past-the-end element in the list container.
func (ll *LinkedList) End() *Iterator {
	return endFwdIterator
}

// RBegin returns a reverse iterator pointing to the last element in the container (i.e., its reverse beginning).
// Reverse iterators iterate backwards: increasing them moves them towards the beginning of the container.
func (ll *LinkedList) RBegin() *Iterator {
	tail := ll.tail.Load().(*Node)
	return &Iterator{currentNode: tail, direction: backwardDirection, index: ll.size}
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

	if ll.size == 0 {
		ll.head.Store(node)
		ll.tail.Store(node)
	} else {
		tmpNode := ll.head.Load().(*Node)
		tmpNode.previous = node
		node.next = tmpNode
		ll.head.Store(node)
	}
	atomic.AddInt64(&ll.size, 1)
}

// PushBack adds a new element at the end of the list container, after its current last element.
// This effectively increases the container size by one.
// Panics if an invalid type is provided.
func (ll *LinkedList) PushBack(val interface{}) {
	element := ll.valueType.Validate(val)
	node := &Node{Value: element}

	if ll.size == 0 {
		ll.head.Store(node)
		ll.tail.Store(node)
	} else {
		tail := ll.tail.Load().(*Node)
		tail.next = node
		node.previous = tail
		ll.tail.Store(node)
	}
	atomic.AddInt64(&ll.size, 1)
}

// PopFront deletes the first element of the list and returns it's value
func (ll *LinkedList) PopFront() interface{} {
	if ll.size == 0 {
		return nil
	}
	head := ll.head.Load().(*Node)
	next := head.next
	next.previous = nil
	ll.head.Store(next)
	atomic.AddInt64(&ll.size, -1)
	return containers.CleanBasicType(head.Value)
}

// PopBack deletes the last element of the list and returns it's value
func (ll *LinkedList) PopBack() interface{} {
	if ll.size == 0 {
		return nil
	}
	tail := ll.tail.Load().(*Node)
	previous := tail.previous
	previous.next = nil
	ll.tail.Store(previous)
	atomic.AddInt64(&ll.size, -1)
	return containers.CleanBasicType(tail.Value)
}

// Clear Removes all elements from the list container (which are destroyed), and leaving the list with a size of 0.
func (ll *LinkedList) Clear() {
	ll.head.Store(nil)
	ll.tail.Store(nil)
	atomic.StoreInt64(&ll.size, 0)
}

/** Capacity Functions **/

// Size returns the length of the linked list
func (ll *LinkedList) Size() int64 {
	return ll.size
}

// Empty returns whether the list container is empty (i.e. whether its size is 0).
func (ll *LinkedList) Empty() bool {
	if ll.size == 0 {
		return true
	}
	return false
}

/** Element Access **/

// Front returns the value of the first element of the linked list
func (ll *LinkedList) Front() interface{} {
	head := ll.head.Load().(*Node)
	return containers.CleanBasicType(head.Value)
}

// Back returns the value of the last element of the linked list
func (ll *LinkedList) Back() interface{} {
	tail := ll.tail.Load().(*Node)
	return containers.CleanBasicType(tail.Value)
}

/** Constructors **/

// New constructs an empty container linked list, with no elements.
func New(valueType containers.Container) *LinkedList {
	return &LinkedList{valueType: valueType}
}

// NewInt constructs an empty integer linked list, with no elements.
func NewInt() *LinkedList {
	return New(containers.IntContainer(0))
}
