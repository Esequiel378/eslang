package stack

import (
	"fmt"
)

// Stack struct    represents a stack of elements.
type Stack struct {
	stack []StackValue
}

// NewStack function    returns a new stack.
func NewStack() Stack {
	return Stack{
		stack: []StackValue{},
	}
}

// IsEmpty method    returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

// Push method    pushes the element onto the stack.
func (s *Stack) Push(value StackValue) {
	s.stack = append(s.stack, value)
}

// Peek method    returns the top element of the stack without removing it.
func (s *Stack) Peek() (StackValue, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Stack.Peek()`, stack is empty")
	}

	// Get the index of the top most element.
	index := len(s.stack) - 1
	// Index into the slice and obtain the element.
	value := (s.stack)[index]

	return value, nil
}

// Pop method    removes and returns the top element of the stack.
func (s *Stack) Pop() (StackValue, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Stack.Pop()`, stack is empty")
	}

	// Get the index of the top most element.
	index := len(s.stack) - 1
	// Index into the slice and obtain the element.
	value := (s.stack)[index]
	// Remove it from the stack by slicing it off.
	s.stack = (s.stack)[:index]

	return value, nil
}

// PopTwo method    removes and returns the top two elements of the stack.
func (s *Stack) PopTwo() (lhs StackValue, rhs StackValue, err error) {
	rhs, err = s.Pop()

	if err != nil {
		return nil, nil, err
	}

	lhs, err = s.Pop()

	if err != nil {
		return nil, nil, err
	}

	return lhs, rhs, nil
}
