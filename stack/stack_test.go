package stack

import "testing"

func TestNew(t *testing.T) {
	s := New(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("Got %d, expected 3", s.Size())
	}
}

func TestStack_Push(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(4)
	if s.Size() != 2 {
		t.Errorf("Got %d, expected 2", s.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	s := New()
	nullVal := s.Pop()
	if nullVal != nil {
		t.Errorf("Got %d, expected nil", nullVal)
	}

	s.Push(1)
	s.Push(2)
	val := s.Pop()
	if val != 2 {
		t.Errorf("Got %d, expected 2", val)
	}
	if s.Size() != 1 {
		t.Errorf("Got %d, expected 1", s.Size())
	}
}

func TestStack_Empty(t *testing.T) {
	s := New()
	if !s.Empty() {
		t.Errorf("Got false, expected true")
	}
	s.Push(1)
	if s.Empty() {
		t.Errorf("Got true, expected false")
	}
}

func TestStack_Peek(t *testing.T) {
	s := New()
	if s.Peek() != nil {
		t.Errorf("Received a value, expected nil")
	}
	s.Push(1)
	s.Push(2)
	if s.Peek() != 2 {
		t.Errorf("Got %d, expected 2", s.Peek())
	}
}
