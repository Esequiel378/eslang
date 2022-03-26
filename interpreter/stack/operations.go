package stack

import (
	"eslang/core"
	"fmt"
)

// TODO: use a generic function for opertaions with numbers

// AddValues function    adds two values together
func AddValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not add values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) + rhs.Value().(int64)), nil
	}

	return nil, fmt.Errorf("can not add values of type %s", lhs.Type())
}

// SubtractValues function    subtracts two values
func SubtractValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not subtract values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) - rhs.Value().(int64)), nil
	}

	return nil, fmt.Errorf("can not subtract values of type %s", lhs.Type())
}

// MultiplyValues function    multiplies two values
func MultiplyValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not multiply values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) * rhs.Value().(int64)), nil
	}

	return nil, fmt.Errorf("can not multiply values of type %s", lhs.Type())
}

// CompareEqualValues function    compares if two values are equal
func CompareEqualValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		if lhs.Value().(int) == rhs.Value().(int) {
			return NewStackValueInt(1), nil
		}
		return NewStackValueInt(0), nil
	}

	return nil, fmt.Errorf("can not compare values of type %s", lhs.Type())
}
