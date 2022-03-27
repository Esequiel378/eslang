package interpreter

import (
	"eslang/core"
	s "eslang/interpreter/stack"
	"fmt"
)

type (
	OPHandler func(*s.Stack, core.Operation) error
)

var REGISTERED_OPERATIONS = map[core.OPType]OPHandler{
	core.OP_DUMP:        OPDump,
	core.OP_PUSH_FLOAT:  OPPushFloat,
	core.OP_PUSH_INT:    OPPushInt,
	core.OP_PUSH_STRING: OPPushStr,
}

func OPPushFloat(stack *s.Stack, _op core.Operation) error {
	op, ok := _op.(core.OperationPushFloat)

	if !ok {
		panic("OPPushFloat: invalid operation type")
	}

	sValue := s.NewStackValueFloat(op.Value())
	stack.Push(sValue)

	return nil
}

func OPPushInt(stack *s.Stack, _op core.Operation) error {
	op, ok := _op.(core.OperationPushInt)

	if !ok {
		panic("OPPushInt: invalid operation type")
	}

	sValue := s.NewStackValueInt(op.Value())
	stack.Push(sValue)

	return nil
}

func OPPushStr(stack *s.Stack, _op core.Operation) error {
	op, ok := _op.(core.OperationPushString)

	if !ok {
		panic("OPPushStr: invalid operation type")
	}

	sValue := s.NewStackValueString(op.Value())
	stack.Push(sValue)

	return nil
}

func OPDump(stack *s.Stack, _ core.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	fmt.Println(sValue.Value())

	return nil
}

func OPDup(stack *s.Stack, _ *core.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	stack.Push(sValue)
	stack.Push(sValue)

	return nil
}
