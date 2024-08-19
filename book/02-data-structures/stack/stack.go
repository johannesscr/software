package stack

import "fmt"

type Stack struct {
	data []any
}

func (s *Stack) String() string {
	return fmt.Sprintf("Stack%v", s.data)
}

// Push inserts a new element onto the stack
func (s *Stack) Push(v any) {
	s.data = append(s.data, v)
}

// IsEmpty check whether the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Top returns the top value in the stack, if the stack is empty it'll return
// nil.
func (s *Stack) Top() any {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// Pop deletes an element from the top of the stack. Returns true if successful.
func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.data = s.data[:len(s.data)-1]
	return true
}

// PopValue deletes an element from the top of the stack and returns the value
// that has been deleted.
func (s *Stack) PopValue() any {
	if s.IsEmpty() {
		return nil
	}
	v := s.Top()
	s.Pop()
	return v
}
