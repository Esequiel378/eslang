package interpreter

import (
	"eslang/core"
	"fmt"
)

type OPHandler func(*Stack, *core.Operation) error

var REGISTERED_OPERATIONS = map[core.OperationType]OPHandler{
	core.OP_PUSH_INT:   OPPushInt,
	core.OP_PUSH_FLOAT: OPPushFloat,
	core.OP_PUSH_STR:   OPPushStr,
	core.OP_MOP:        OPMop,
	core.OP_DUMP:       OPDump,
	core.OP_VAR:        OPMem,
}

func OPPushFloat(stack *Stack, op *core.Operation) error {
	sValue := NewStackValue().SetFloat(op.Value().Float())
	stack.Push(sValue)

	return nil
}

func OPPushInt(stack *Stack, op *core.Operation) error {
	sValue := NewStackValue().SetInt(op.Value().Int())
	stack.Push(sValue)

	return nil
}

func OPPushStr(stack *Stack, op *core.Operation) error {
	sValue := NewStackValue().SetStr(op.Value().Str())
	stack.Push(sValue)

	return nil
}

func OPDump(stack *Stack, _ *core.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	v, err := sValue.Value()
	if err != nil {
		return err
	}

	fmt.Println(v)

	return nil
}

var REGISTERED_MOPS = map[core.TokenType]func(*Stack, *core.Operation) error{
	core.TOKEN_EQUAL: OPEqual,
	core.TOKEN_MINUS: OPMinus,
	core.TOKEN_PLUS:  OPPlus,
}

func OPMop(stack *Stack, op *core.Operation) error {
	handler, ok := REGISTERED_MOPS[op.TokenStart().Token()]

	if !ok {
		return fmt.Errorf("exaustive MOPs handiling. %s not found", op.TokenEnd().TokenAlias())
	}

	if err := handler(stack, op); err != nil {
		return err
	}

	return nil
}

func OPPlus(stack *Stack, _ *core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	if !lhs.IsNumber() || !rhs.IsNumber() {
		return fmt.Errorf("can not add `%s` with `%s`", lhs.TypeAlias(), rhs.TypeAlias())
	}

	normalizeNumbers(lhs, rhs)

	sValue := NewStackValue()

	switch lhs.Type() {
	case core.Int:
		sValue.SetInt(lhs.Int() + rhs.Int())
	case core.Float:
		sValue.SetFloat(lhs.Float() + rhs.Float())
	}

	stack.Push(sValue)

	return nil
}

func OPMinus(stack *Stack, _ *core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	normalizeNumbers(lhs, rhs)

	sValue := NewStackValue()

	switch lhs.Type() {
	case core.Int:
		sValue.SetInt(lhs.Int() - rhs.Int())
	case core.Float:
		sValue.SetFloat(lhs.Float() - rhs.Float())
	}

	stack.Push(sValue)

	return nil
}

func OPEqual(stack *Stack, _ *core.Operation) error {
	_lhs, _rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	lhs, err := _lhs.Value()
	if err != nil {
		return err
	}

	rhs, err := _rhs.Value()
	if err != nil {
		return err
	}

	sValue := NewStackValue()

	// TODO: at some point this should be using bool type
	if lhs == rhs {
		sValue.SetInt(1)
	}

	stack.Push(sValue)

	return nil
}

func OPMem(stack *Stack, op *core.Operation) error {
	opValue := op.Value()

	sValue := NewStackValue().SetType(opValue.Type())

	stack.SetVariable("temp", sValue)

	return nil
}
