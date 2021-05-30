package bit

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	_ = NewTree([]int{1, 10, 3, 6})
}

func TestUpdate(t *testing.T) {
	bit := NewTree([]int{1, 10, 3, 6})
	bit.Update(1, 2)
}

func TestGetSum(t *testing.T) {
	bit := NewTree([]int{1, 10, 3, 6})
	bit.GetSum(1)
}

func TestGetRange(t *testing.T) {
	bit := NewTree([]int{1, 10, 3, 6})
	bit.GetRange(4, 2)
}
