// Stack implemented by slice
package data_structure

type stack[T any] struct {
	vals []T
}

// NewStack get a stack
func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

// Push push a val into stack
func (s *stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

// Pop pop a val from stack
func (s *stack[T]) Pop() T {
	lastIndex := s.Len() - 1
	val := s.vals[lastIndex]
	s.vals = s.vals[0:lastIndex]
	return val
}

// Len get stack's length
func (s *stack[T]) Len() int {
	return len(s.vals)
}

// IsEmpty tell you if the stack is empty
func (s *stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Clear empty stack
func (s *stack[T]) Clear() {
	s.vals = s.vals[0:0]
}
