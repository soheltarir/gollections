package bit

import "errors"

type bit []int

func (b bit) Len() int {
	return len(b)
}

func (b bit) Update(index int, value int) {
	index += 1
	for index <= b.Len() {
		b[index] += value
		index += index & -index
	}
}

func (b bit) GetSum(index int) int {
	res := 0
	index += 1
	for index > 0 {
		res += b[index]
		index -= index & -index
	}
	return res
}

func (b bit) GetRange(index1 int, index2 int) (int, error) {
	if index1 > index2 {
		err := errors.New("value1 cannot be more than value2")
		return 0, err
	}
	return b.GetSum(index2) - b.GetSum(index1), nil
}

func NewTree(array []int) bit {
	var bit bit
	initBit := make([]int, len(array)+1)
	bit = append(bit, initBit...)
	for index, element := range array {
		bit.Update(index, element)
	}
	return bit
}
