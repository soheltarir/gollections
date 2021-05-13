package queue

import "fmt"

func Example() {
	// Create a new queue with some initial data
	queue := NewString("a", "b", "c")

	// Retrieve the next element in the queue
	fmt.Println(queue.Dequeue())

	// Add new element to the queue
	queue.Enqueue("d")

	// Inspect the next element in the queue without popping it
	fmt.Println(queue.Front())

	// Inspect the last added element in the queue
	fmt.Println(queue.Back())

	// Clear the entire queue
	queue.Clear()

	// Output:
	// a
	// b
	// d
}
