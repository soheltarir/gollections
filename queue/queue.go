package queue

// A Queue is a linear structure which follows a particular order in which the operations are performed.
// The order is First In First Out (FIFO)
type Queue struct {
	data	[]interface{}
	size	int
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(value interface{}) {
	q.data = append(q.data, value)
	q.size++
}

// Dequeue removes an item from the queue. The items are popped in the same order in which they are pushed.
func (q *Queue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	value := q.data[0]
	q.data = q.data[1:]
	q.size--
	return value
}

// Size returns the total size of the queue
func (q *Queue) Size() int {
	return q.size
}

// Empty returns true if the queue has no items
func (q *Queue) Empty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

// New instantiates a new queue with the items provided (order is preserved)
func New(values ...interface{}) *Queue {
	q := &Queue{}
	for _, value := range values {
		q.Enqueue(value)
	}
	return q
}
