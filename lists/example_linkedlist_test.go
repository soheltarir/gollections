// Lists are sequence containers that allow constant time insert and erase operations anywhere within the sequence,
// and iteration in both directions.
// List containers are implemented as doubly-linked lists; Doubly linked lists can store each of the elements they
// contain in different and unrelated storage locations. The ordering is kept internally by the association to each
// element of a link to the element preceding it and a link to the element following it.
package lists

import "fmt"

func Example() {
	// Instantiates a int list
	list := NewInt()

	// Insert elements. list.Begin() points to head of the list
	list.Insert(list.Begin(), 1, 2, 3)
	// Prints the list
	fmt.Println(list.Display())

	// Push at the back of the list
	list.PushBack(4)
	fmt.Println(list.Display())

	// Add new element at the front of the list
	list.PushFront(5)
	fmt.Println(list.Display())

	fmt.Println(list.PopFront()) // Pops the head of the list
	fmt.Println(list.PopBack())  // Pops the tail of the list

	// Initialise an iterator from the front of the list
	it := list.Begin()
	// Advance the iterator by 1 step
	it.Advance(1)

	// Erase the element at the position of the iterator
	list.Erase(it)
	fmt.Println(list.Display())

	// Output:
	// 1 <-> 2 <-> 3
	// 1 <-> 2 <-> 3 <-> 4
	// 5 <-> 1 <-> 2 <-> 3 <-> 4
	// 5
	// 4
	// 1 <-> 3
}

func Example_iteration() {
	// Create a New list
	list := NewInt()
	list.Insert(list.Begin(), 1, 2, 3, 4, 5)

	// Iterate through the list in forward direction
	for it := list.Begin(); it != list.End(); it = it.Next() {
		fmt.Printf("%d->", it.Value())
	}
	fmt.Printf("\n")

	// Iterate through the list in backward direction
	for it := list.RBegin(); it != list.REnd(); it = it.Next() {
		fmt.Printf("%d->", it.Value())
	}
	fmt.Printf("\n")

	// Output:
	// 1->2->3->4->5->
	// 5->4->3->2->1->
}
