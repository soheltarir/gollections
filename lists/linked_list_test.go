package lists

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
	"github.com/stretchr/testify/assert"
	"testing"
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
	fmt.Println("head: ", ll.head)
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
