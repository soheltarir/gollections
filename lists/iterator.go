// Package lists Lists are sequence containers that allow constant time insert and erase operations anywhere within the sequence,
// and iteration in both directions.
// List containers are implemented as doubly-linked lists; Doubly linked lists can store each of the elements they
// contain in different and unrelated storage locations. The ordering is kept internally by the association to each
// element of a link to the element preceding it and a link to the element following it.
package lists

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
)

type direction uint

const (
	forwardDirection direction = iota
	backwardDirection
)

// Iterator is a stateful iterator for traversing a linked list.
// Please note, that iteration over a list is not a thread-safe operation, and if parallel write operations are
// being performed on the list, the traversal can provide stale and outdated data.
type Iterator struct {
	currentNode *Node
	direction   direction
	index       int64
	limit       int64
}

// Next returns the iterator to the next/previous element in the list based on the traversal direction. Panics if
// the iterator reaches out of bounds
// Please note, iteration over a list is not a thread-safe operation.
func (it *Iterator) Next() *Iterator {
	currNode := it.currentNode
	if currNode == nil {
		panic("iterator crossed list's bounds")
	}
	nextIt := &Iterator{limit: it.limit, direction: it.direction}
	switch it.direction {
	case forwardDirection:
		if it.currentNode.next == nil {
			nextIt = endFwdIterator
		} else {
			nextIt.currentNode, nextIt.index = it.currentNode.next, it.index+1
		}
	case backwardDirection:
		if it.currentNode.previous == nil {
			nextIt = endBackIterator
		} else {
			nextIt.currentNode, nextIt.index = it.currentNode.previous, it.index-1
		}
	}
	return nextIt
}

func (it *Iterator) IsEqual(it2 *Iterator) bool {
	if it.currentNode == it2.currentNode {
		return true
	}
	return false
}

// Advance moves the iterator forward by the no. of the steps provided
func (it *Iterator) Advance(steps int) (*Iterator, error) {
	if steps <= 0 {
		return it, fmt.Errorf("step size should be greater than zero")
	}
	newItr := it
	for steps > 0 {
		newItr = newItr.Next()
		steps--
	}
	*it = *newItr
	return it, nil
}

// Value returns the current element's value
func (it *Iterator) Value() interface{} {
	return containers.CleanBasicType(it.currentNode.Value)
}

// Index returns the current element's index in the list
func (it *Iterator) Index() int64 {
	return it.index
}

// endFwdIterator is used for signifying the end of an iteration
var endFwdIterator = &Iterator{currentNode: nil, direction: forwardDirection}

// endBackIterator is used for signifying the end of an iteration in reverse direction
var endBackIterator = &Iterator{currentNode: nil, direction: backwardDirection}
