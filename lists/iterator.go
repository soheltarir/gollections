package lists

import "github.com/soheltarir/gollections/containers"

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
}

// Next returns the iterator to the next/previous element in the list based on the traversal direction
// Please note, iteration over a list is not a thread-safe operation.
func (it *Iterator) Next() *Iterator {
	currNode := it.currentNode
	switch it.direction {
	case forwardDirection:
		it.currentNode = currNode.next
		if it.currentNode == nil {
			it = endFwdIterator
		} else {
			it.index++
		}
	case backwardDirection:
		it.currentNode = currNode.previous
		if it.currentNode == nil {
			it = endBackIterator
		} else {
			it.index--
		}
	}
	return it
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
