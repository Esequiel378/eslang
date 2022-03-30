package stack

import (
	"eslang/core"
	"fmt"
)

// StackValue interface    represents a value in the stack.
type StackValue interface {
	Type() core.Type
	Value() any
	TestTruthy() (bool, error)
}

// Stack struct    represents a stack of elements.
type Stack struct {
	stack     []StackValue
	variables map[string]StackValue
}

// NewStack function    returns a new stack.
func NewStack() Stack {
	return Stack{
		stack:     []StackValue{},
		variables: make(map[string]StackValue),
	}
}

// GetVariable method    returns a StackValue for the variable and a boolean indicating if the variable was found.
func (s *Stack) GetVariable(name string) (StackValue, bool) {
	value, found := s.variables[name]

	return value, found
}

// SetVariable method    sets the value of the variable.
func (s *Stack) SetVariable(name string, value StackValue) {
	s.variables[name] = value
}

// IsEmpty method    returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

// Push method    pushes the element onto the stack.
func (s *Stack) Push(value StackValue) {
	s.stack = append(s.stack, value)
}

// PeekAt method    returns the element at the specified index.
// Accepts negative indices, which are relative to the end of the stack.
func (s *Stack) PeekAt(index int) (StackValue, error) {
	if index < 0 {
		index = len(s.stack) + index
	}

	if index < 0 || index >= len(s.stack) {
		return nil, fmt.Errorf("index out of bounds")
	}

	return s.stack[index], nil
}

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
// TODO: Improve this method, it is not very efficient.
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
