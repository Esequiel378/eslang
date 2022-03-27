package interpreter

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

type (
	OPHandler func(*s.Stack, ops.Operation) error
)

var REGISTERED_OPERATIONS = map[ops.OPType]OPHandler{
	ops.OP_DUMP:        OPDump,
	ops.OP_PUSH_FLOAT:  OPPushFloat,
	ops.OP_PUSH_INT:    OPPushInt,
	ops.OP_PUSH_STRING: OPPushStr,
}

func OPPushFloat(stack *s.Stack, _op ops.Operation) error {
	op, ok := _op.(ops.OPPushFloat)

	if !ok {
		panic("OPPushFloat: invalid operation type")
	}

	sValue := s.NewStackValueFloat(op.Value())
	stack.Push(sValue)

	return nil
}

func OPPushInt(stack *s.Stack, _op ops.Operation) error {
	op, ok := _op.(ops.OPPushInt)

	if !ok {
		panic("OPPushInt: invalid operation type")
	}

	sValue := s.NewStackValueInt(op.Value())
	stack.Push(sValue)

	return nil
}

func OPPushStr(stack *s.Stack, _op ops.Operation) error {
	op, ok := _op.(ops.OPPushString)

	if !ok {
		panic("OPPushStr: invalid operation type")
	}

	sValue := s.NewStackValueString(op.Value())
	stack.Push(sValue)

	return nil
}

func OPDump(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	fmt.Println(sValue.Value())

	return nil
}

func OPDup(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	stack.Push(sValue)
	stack.Push(sValue)

	return nil
}
