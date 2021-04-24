package stack

type Stack struct {
	data	[]interface{}
	size	int
}

// New instantiates a fresh stack with the values provided
func New(values ...interface{}) *Stack {
	obj := &Stack{data: make([]interface{}, 0)}
	obj.data = append(obj.data, values...)
	obj.size = len(values)
	return obj
}

// Push insert element
func (s *Stack) Push(value interface{}) {
	s.data = append(s.data, value)
	s.size++
}

// Pop removes the top element
func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return value
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return s.size
}

// Empty returns whether the stack is empty
func (s *Stack) Empty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

// Peek returns the value of the next element
func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.data[s.size-1]
}
