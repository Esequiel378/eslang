package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

func OPPushFloat(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushFloat)

	sValue := s.NewStackValueFloat(op.Value())
	stack.Push(sValue)

	return nil
}

func OPPushInt(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushInt)

	sValue := s.NewStackValueInt(op.Value())
	stack.Push(sValue)

	return nil
}

func OPPushStr(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushString)

	sValue := s.NewStackValueString(op.Value())
	stack.Push(sValue)

	return nil
}
