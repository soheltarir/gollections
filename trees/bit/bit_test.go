package bit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	bit := NewTree([]int{1, 10, 3, 6})
	assert.IsType(t, []int{}, bit.array)
}

func TestUpdate(t *testing.T) {
	nums := []int{1, 10, 3, 6}
	bit := NewTree(nums)
	assert.IsType(t, []int{}, bit.array)
	oldValue := bit.array[2]
	fmt.Println(bit.array)
	nums[1] += 2
	bit.Update(1, 2)
	fmt.Println(bit.array)
	assert.Equal(t, oldValue+2, bit.array[2])
}

func TestGetSum(t *testing.T) {
	bit := NewTree([]int{1, 10, 3, 6})
	assert.IsType(t, []int{}, bit.array)
	sum := bit.GetSum(2)
	assert.Equal(t, 14, sum)
}

func TestGetRange(t *testing.T) {
	bit := NewTree([]int{1, 10, 3, 6})
	assert.IsType(t, []int{}, bit.array)
	var rangeSum int
	rangeSum, _ = bit.GetRange(2, 1)
	assert.Equal(t, 0, rangeSum)
	rangeSum, _ = bit.GetRange(1, 2)
	assert.Equal(t, 13, rangeSum)
}
