package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	s := NewInt(1, 2, 3)
	assert.Equal(t, int64(3), s.Size())
}

func TestStack_Push(t *testing.T) {
	s := NewInt()
	s.Push(1)
	assert.Equal(t, 1, s.data.Back())
	s.Push(4)
	assert.Equal(t, 4, s.data.Back())
}

func TestStack_Pop(t *testing.T) {
	s := NewInt()
	assert.Nil(t, s.Pop())

	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, int64(1), s.Size())
}

func TestStack_Empty(t *testing.T) {
	s := NewInt()
	assert.True(t, s.Empty())
	s.Push(1)
	assert.False(t, s.Empty())
}

func TestStack_Top(t *testing.T) {
	s := NewString()
	assert.Nil(t, s.Top())
	s.Push("a")
	s.Push("b")
	assert.Equal(t, "b", s.Top())
}

func TestStack_Clear(t *testing.T) {
	s := NewString("a", "b", "c")
	assert.Greater(t, s.Size(), int64(0))
	s.Clear()
	assert.Zero(t, s.Size())
}

func Example() {
	// Create a new stack
	stack := NewString("a", "b", "c")

	// Retrieve the next element in stack
	fmt.Println(stack.Pop())

	// Check the value of the next element in stack
	fmt.Println(stack.Top())

	// Add a new element in the stack
	stack.Push("d")
	fmt.Println(stack.Top())

	// Check the size of the stack
	fmt.Println(stack.Size())

	// Clear the stack
	stack.Clear()

	// Output:
	// c
	// b
	// d
	// 3
}
