package lists

import (
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
	ll := NewInt()
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushFront(3)
	ll.PushFront(4)
	var result []interface{}
	for it := ll.Begin(); it != ll.End(); it = it.Next() {
		result = append(result, it.Value())
	}
	assert.Equal(t, []interface{}{4, 3, 1, 2}, result)
}

func TestListBackwardIteration(t *testing.T) {
	ll := NewInt()
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushFront(3)
	ll.PushFront(4)
	var result []interface{}
	for it := ll.RBegin(); it != ll.REnd(); it = it.Next() {
		result = append(result, it.Value())
	}
	assert.Equal(t, []interface{}{2, 1, 3, 4}, result)
}
