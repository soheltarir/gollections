package lists

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	list1 := New()
	list2 := &LinkedList{}
	if reflect.TypeOf(list1) != reflect.TypeOf(list2) {
		t.Errorf("Got %T, expected %T", list1, list2)
	}
}

func TestLinkedList_Add(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	if list.Size() != 2 {
		t.Errorf("Got %d expected 2", list.Size())
	}
	if list.head.Value != 1 {
		t.Errorf("Got %d expected 1", list.head.Value)
	}
	if list.tail.Value != 2 {
		t.Errorf("Got %d expected 2", list.head.Value)
	}
}

func TestLinkedList_Size(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	if list.Size() != 2 {
		t.Errorf("Got %d expected 2", list.Size())
	}
}

func TestLinkedList_Clear(t *testing.T) {
	list := New()
	list.Add(1)
	list.Clear()
	if list.Size() > 0 {
		t.Errorf("Got %d expected 0", list.Size())
	}
}

func TestLinkedList_Back(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.Back() != 3 {
		t.Errorf("Got %d expected 3", list.Back())
	}
}

func TestLinkedList_Front(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.Front() != 1 {
		t.Errorf("Got %d expected 3", list.Front())
	}
}

func TestLinkedList_PopBack(t *testing.T) {
	list := New()
	nullVal := list.PopBack()
	if nullVal != nil {
		t.Errorf("Got %d expected nil", nullVal)
	}
	list.Add(1)
	list.Add(2)
	val := list.PopBack()
	if val != 2 {
		t.Errorf("Got %d expected 2", val)
	}
	if list.Size() != 1 {
		t.Errorf("Got %d expected 1", list.Size())
	}
}

func TestLinkedList_PopFront(t *testing.T) {
	list := New()
	nullVal := list.PopFront()
	if nullVal != nil {
		t.Errorf("Got %d expected nil", nullVal)
	}
	list.Add(1)
	list.Add(2)
	val := list.PopFront()
	if val != 1 {
		t.Errorf("Got %d expected 1", val)
	}
	if list.Size() != 1 {
		t.Errorf("Got %d expected 1", list.Size())
	}
}
