package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

// OPPushFloat function    pushes a float value onto the stack.
func OPPushFloat(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushFloat)

	sValue := s.NewStackValueFloat(op.Value())
	stack.Push(sValue)

	return nil
}

// OPPushInt function    pushes an integer value onto the stack.
func OPPushInt(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushInt)

	sValue := s.NewStackValueInt(op.Value())
	stack.Push(sValue)

	return nil
}

// OPPushStr function    pushes a string value onto the stack.
func OPPushStr(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushString)

	sValue := s.NewStackValueString(op.Value())
	stack.Push(sValue)

	return nil
}

// OPPushBool function    pushes a boolean value onto the stack.
func OPPushBool(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushBool)

	sValue := s.NewStackValueBool(op.Value())
	stack.Push(sValue)

	return nil
}
