package stack

import "eslang/core"

// StackValue interface    represents a value in the stack.
type StackValue interface {
	Type() core.Type
	Value() any
	TestTruthy() (bool, error)
}

// StackValueInt struct    represents an integer value.
type StackValueInt struct {
	value int64
}

// NewStackValueInt function    creates a new integer value.
func NewStackValueInt(value int64) StackValue {
	return StackValueInt{
		value: value,
	}
}

// Type method    returns the type of the value.
func (v StackValueInt) Type() core.Type {
	return core.Int
}

// Value method    returns the value of the value.
func (v StackValueInt) Value() any {
	return v.value
}

// TestTruthy method    returns true if the value is truthy.
func (v StackValueInt) TestTruthy() (bool, error) {
	return v.value != 0, nil
}

// StackValueFloat struct    represents a floating point value.
type StackValueFloat struct {
	value float64
}

// NewStackValueFloat function    creates a new floating point value.
func NewStackValueFloat(value float64) StackValue {
	return StackValueFloat{
		value: value,
	}
}

// Type method    returns the type of the value.
func (v StackValueFloat) Type() core.Type {
	return core.Float
}

// Value method    returns the value of the value.
func (v StackValueFloat) Value() any {
	return v.value
}

// TestTruthy method    returns true if the value is truthy.
func (v StackValueFloat) TestTruthy() (bool, error) {
	return v.value != 0, nil
}

// StackValueString struct    represents a string value.
type StackValueString struct {
	value string
}

// NewStackValueString function    creates a new string value.
func NewStackValueString(value string) StackValue {
	return StackValueString{
		value: value,
	}
}

// Type method    returns the type of the value.
func (v StackValueString) Type() core.Type {
	return core.String
}

// Value method    returns the value of the value.
func (v StackValueString) Value() any {
	return v.value
}

// TestTruthy method    returns true if the value is truthy.
func (v StackValueString) TestTruthy() (bool, error) {
	return v.value != "", nil
}
