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
