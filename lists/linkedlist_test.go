// Lists are sequence containers that allow constant time insert and erase operations anywhere within the sequence,
// and iteration in both directions.
// List containers are implemented as doubly-linked lists; Doubly linked lists can store each of the elements they
// contain in different and unrelated storage locations. The ordering is kept internally by the association to each
// element of a link to the element preceding it and a link to the element following it.
package lists

import (
	"math/rand"
	"testing"
	"time"

	"github.com/soheltarir/gollections/containers"
	"github.com/stretchr/testify/assert"
)

func TestNewInt(t *testing.T) {
	actual := NewInt()
	expected := &LinkedList{valueType: containers.IntContainer(0)}
	assert.IsType(t, expected, actual)
}

func TestLinkedList_PushFront(t *testing.T) {
	ll := NewInt()
	ll.PushFront(1)
	ll.PushFront(2)
	assert.Equal(t, 2, ll.Front())
}

func TestLinkedList_PushBack(t *testing.T) {
	ll := NewInt()
	ll.PushBack(1)
	ll.PushFront(2)
	assert.Equal(t, 1, ll.Back())
}

func TestListForwardIteration(t *testing.T) {
	var result []interface{}
	ll := NewInt()

	// Test iteration on empty list
	for it := ll.Begin(); it != ll.End(); it = it.Next() {
		result = append(result, it.Value())
	}
	assert.Empty(t, result)

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushFront(3)
	ll.PushFront(4)
	// Empty the result
	result = make([]interface{}, 0)
	var lastIndex int64
	for it := ll.Begin(); it != ll.End(); it = it.Next() {
		result = append(result, it.Value())
		lastIndex = it.index
	}
	assert.Equal(t, []interface{}{4, 3, 1, 2}, result)
	// Test Last index
	assert.Equal(t, int64(3), lastIndex)
}

func TestListBackwardIteration(t *testing.T) {
	var result []interface{}
	ll := NewInt()

	// Test iteration on empty list
	for it := ll.RBegin(); it != ll.REnd(); it = it.Next() {
		result = append(result, it.Value())
	}
	assert.Empty(t, result)

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushFront(3)
	ll.PushFront(4)
	var lastIndex int64
	for it := ll.RBegin(); it != ll.REnd(); it = it.Next() {
		result = append(result, it.Value())
		lastIndex = it.Index()
	}
	assert.Equal(t, []interface{}{2, 1, 3, 4}, result)
	assert.Equal(t, int64(0), lastIndex)
}

func TestLinkedList_PopFront(t *testing.T) {
	ll := NewInt()
	// Test empty result
	assert.Nil(t, ll.PopFront())
	ll.PushBack(1)
	ll.PushBack(2)
	assert.Equal(t, 1, ll.PopFront())
	assert.Equal(t, int64(1), ll.Size())
}

func TestLinkedList_PopBack(t *testing.T) {
	ll := NewInt()
	// Test empty result
	assert.Nil(t, ll.PopBack())
	ll.PushBack(1)
	ll.PushBack(2)
	assert.Equal(t, 2, ll.PopBack())
	assert.Equal(t, int64(1), ll.Size())
}

func TestLinkedList_Empty(t *testing.T) {
	ll := NewInt()
	assert.True(t, ll.Empty())
	ll.PushBack(1)
	assert.False(t, ll.Empty())
}

func TestLinkedList_Clear(t *testing.T) {
	ll := NewInt()
	ll.PushFront(1)
	ll.PushBack(2)
	ll.Clear()
	assert.True(t, ll.Empty())
}

func TestLinkedList_Front(t *testing.T) {
	ll := NewInt()
	assert.Nil(t, ll.Front())

	ll.PushFront(1)
	assert.Equal(t, 1, ll.Front())
}

func TestLinkedList_Back(t *testing.T) {
	ll := NewInt()
	assert.Nil(t, ll.Back())

	ll.PushBack(2)
	assert.Equal(t, 2, ll.Back())
}

func TestLinkedList_Insert(t *testing.T) {
	ll := NewInt()
	ll.Insert(ll.Begin(), 1, 2, 3)
	assert.Equal(t, 1, ll.Front())
	// Insert between
	it := ll.Begin()
	it.Advance(1)
	ll.Insert(it, 4, 5)
	var actual []interface{}
	for i := ll.Begin(); i != ll.End(); i = i.Next() {
		actual = append(actual, i.Value())
	}
	expected := []interface{}{1, 4, 5, 2, 3}
	assert.Equal(t, expected, actual)
}

func TestLinkedList_Erase(t *testing.T) {
	ll := NewInt()
	ll.Insert(ll.Begin(), 1, 2, 3)
	it := ll.Begin()
	_, _ = it.Advance(1)
	assert.NoError(t, ll.Erase(it))

	// Test Single erase
	var actual []interface{}
	for i := ll.Begin(); i != ll.End(); i = i.Next() {
		actual = append(actual, i.Value())
	}
	assert.Equal(t, []interface{}{1, 3}, actual)

	// Test Range erase
	ll.Insert(ll.Begin(), 4, 5, 6, 7) // 4, 5, 6, 7, 1, 3
	it1, it2 := ll.Begin(), ll.Begin()
	_, _ = it2.Advance(4)
	assert.NoError(t, ll.Erase(it1, it2))

	actual = make([]interface{}, 0)
	for i := ll.Begin(); i != ll.End(); i = i.Next() {
		actual = append(actual, i.Value())
	}
	assert.Equal(t, []interface{}{1, 3}, actual)

	// Handle error
	assert.Error(t, ll.Erase(it, it1, it2))

	// Handle deletion of tail
	assert.NoError(t, ll.Erase(ll.RBegin()))

}

func TestIterator_Next(t *testing.T) {
	ll := NewInt()
	ll.Insert(ll.Begin(), 1, 2)

	it := ll.Begin()
	it = it.Next()
	assert.Equal(t, 2, it.Value())

	assert.Panics(t, func() {
		it = it.Next()
		it = it.Next()
	})

	it = ll.RBegin()
	it = it.Next()
	assert.Equal(t, 1, it.Value())
	assert.Panics(t, func() {
		it = it.Next()
		it = it.Next()
	})
}

func TestIterator_Advance(t *testing.T) {
	ll := NewInt()
	ll.Insert(ll.Begin(), 1, 2, 3)

	it := ll.Begin()
	_, err := it.Advance(-1)
	assert.Error(t, err)
}

func TestLinkedList_Display(t *testing.T) {
	ll := NewInt()
	ll.Insert(ll.Begin(), 1, 2, 3, 4, 5)
	expected := "1 <-> 2 <-> 3 <-> 4 <-> 5"
	assert.Equal(t, expected, ll.Display())
}

func TestLinkedList_Search(t *testing.T) {
	ll := NewInt()
	ll.Insert(ll.Begin(), 1, 2, 3, 4, 5, 6)
	assert.Equal(t, int64(2), ll.Search(3).Index())
	assert.Nil(t, ll.Search(7))
}

func BenchmarkLinkedList_Search(b *testing.B) {
	ll := NewInt()
	slice := make([]interface{}, 10000, 10000)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	ll.Insert(ll.Begin(), slice...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Search(rand.Intn(999))
	}
}
