package stack

import (
	"eslang/core"
	"fmt"
)

// StackValue interface    Represents a value in the stack.
type StackValue interface {
	Type() core.Type
	Value() any
	TestTruthy() (bool, error)
}

// Stack struct    Represents a stack of elements.
type Stack struct {
	// TODO: considering to use liked list instead of slice.
	stack     []StackValue
	variables map[string]StackValue
}

// NewStack function    Returns a new stack.
func NewStack() Stack {
	// TODO: Set a limit on the stack size. (add a flag to control this)
	return Stack{
		stack:     make([]StackValue, 0),
		variables: make(map[string]StackValue),
	}
}

// Content method    Returns the stack content.
func (s *Stack) Content() []StackValue {
	return s.stack
}

// Size method    Returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.stack)
}

// GetVariable method    Returns a StackValue for the variable and a boolean indicating if the variable was found.
func (s *Stack) GetVariable(name string) (StackValue, bool) {
	value, found := s.variables[name]

	return value, found
}

// SetVariable method    Sets the value of the variable.
func (s *Stack) SetVariable(name string, value StackValue) {
	s.variables[name] = value
}

// IsEmpty method    Returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

// Push method    Pushes the element onto the stack.
func (s *Stack) Push(value StackValue) {
	s.stack = append(s.stack, value)
}

// PeekAt method    Returns the element at the specified index.
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

// PeekTwo method    Returns the top two elements of the stack
func (s *Stack) PeekTwo() (lhs, rhs StackValue, err error) {
	rhs, err = s.Peek()
	if err != nil {
		return lhs, rhs, err
	}

	lhs, err = s.PeekAt(-2)
	if err != nil {
		return lhs, rhs, err
	}

	return lhs, rhs, err
}

// Pop method    Removes and returns the top element of the stack.
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

// PopTwo method    Removes and returns the top two elements of the
// stack as pairs of cells.
func (s *Stack) PopTwo() (lhs, rhs StackValue, err error) {
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
