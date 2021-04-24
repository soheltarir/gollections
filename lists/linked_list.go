package lists

type Node struct {
	Value	interface{}
	next	*Node
	previous *Node
}

type LinkedList struct {
	head	*Node
	tail	*Node
	size	int
}

// Add inserts a new node (with the value) and the end of the linked list
func (ll *LinkedList) Add(value interface{}) {
	newNode := &Node{Value: value}
	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		newNode.previous = ll.tail
		ll.tail = newNode
	}
	ll.size++
}

// Size returns the length of the linked list
func (ll *LinkedList) Size() int {
	return ll.size
}

// Clear removes all elements of the linked list
func (ll *LinkedList) Clear() {
	ll.head, ll.tail = nil, nil
	ll.size = 0
}

// Front returns the value of the first element of the linked list
func (ll *LinkedList) Front() interface{} {
	return ll.head.Value
}

// Back returns the value of the last element of the linked list
func (ll *LinkedList) Back() interface{} {
	return ll.tail.Value
}

// PopFront deletes the first element of the list and returns it's value
func (ll *LinkedList) PopFront() interface{} {
	if ll.size == 0 {
		return nil
	}
	value := ll.head.Value
	ll.head = ll.head.next
	ll.size--
	return value
}

// PopBack deletes the last element of the list and returns it's value
func (ll *LinkedList) PopBack() interface{} {
	if ll.size == 0 {
		return nil
	}
	value := ll.tail.Value
	ll.tail = ll.tail.previous
	ll.tail = nil
	ll.size--
	return value
}

// New instantiates an empty linked list
func New() *LinkedList {
	return &LinkedList{}
}
