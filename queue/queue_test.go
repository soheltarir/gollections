package queue

import "testing"

func TestNew(t *testing.T) {
	q := New(1, 2, 3)
	if q.Size() != 3 {
		t.Errorf("Got %d, expected 3", q.Size())
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	if q.Size() != 1 {
		t.Errorf("Got %d, expected 1", q.Size())
	}
	q.Enqueue(2)
	if q.Size() != 2 {
		t.Errorf("Got %d, expected 2", q.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := New()
	if q.Dequeue() != nil {
		t.Errorf("Got a value, expected nil")
	}
	q.Enqueue(1)
	if q.Dequeue() != 1 {
		t.Errorf("Got unexpected value")
	}
}

func TestQueue_Empty(t *testing.T) {
	q := New()
	if !q.Empty() {
		t.Errorf("Queue should be empty")
	}
	q.Enqueue(10)
	if q.Empty() {
		t.Errorf("Queue is should not be empty")
	}
}
