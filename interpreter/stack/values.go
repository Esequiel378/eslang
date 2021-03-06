package stack

import (
	"eslang/core"
	"fmt"
)

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

// StackValueBool
type StackValueBool struct {
	value bool
}

// NewStackValueBool function    creates a new boolean value.
func NewStackValueBool(value bool) StackValue {
	return StackValueBool{
		value: value,
	}
}

// Type method    return the type of the value.
func (v StackValueBool) Type() core.Type {
	return core.Bool
}

// Value method    returns a boolean value.
func (v StackValueBool) Value() any {
	return v.value
}

// TestTruthy method    returns true if the value is truthy.
func (v StackValueBool) TestTruthy() (bool, error) {
	return v.value, nil
}

// StackValueVar struct    contains a variable value.
type StackValueVar struct {
	value StackValue
	name  string
}

// NewStackValueVariable function    creates a new variable value.
func NewStackValueVariable(name string, value StackValue) StackValueVar {
	return StackValueVar{
		value: value,
		name:  name,
	}
}

// Type method    returns the type of the variable value.
func (v StackValueVar) Type() core.Type {
	if v.value == nil {
		return core.Nil
	}

	return v.value.Type()
}

// Value method    returns the value of the variable.
func (v StackValueVar) Value() any {
	if v.value == nil {
		return nil
	}

	return v.value.Value()
}

// TestTruthy method    returns true if the value of the variable is truthy.
func (v StackValueVar) TestTruthy() (bool, error) {
	if v.value == nil {
		return false, fmt.Errorf("variable %s is not defined", v.name)
	}

	return v.value.TestTruthy()
}

// Name method    returns the name of the variable.
func (v StackValueVar) Name() string {
	return v.name
}
