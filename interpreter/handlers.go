package interpreter

import (
	"eslang/core"
	"fmt"
)

var REGISTERED_OPERATIONS = map[core.OperationType]func(*core.Stack, core.Operation) error{
	core.OP_PUSH:  OPPush,
	core.OP_PLUS:  OPPlus,
	core.OP_MINUS: OPMinus,
	core.OP_EQUAL: OPEqual,
	core.OP_DUMP:  OPDump,
}

func OPPush(stack *core.Stack, op core.Operation) error {
	value := op.Value()

	stack.Push(value)

	return nil
}

func OPPlus(stack *core.Stack, _ core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	lhsf, rhsf, keepFloat, err := normalizeNumbers(lhs, rhs)
	if err != nil {
		return err
	}

	if keepFloat {
		stack.Push(lhsf + rhsf)
	} else {
		stack.Push(int64(lhsf + rhsf))
	}

	return nil
}

func OPMinus(stack *core.Stack, _ core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	lhsf, rhsf, keepFloat, err := normalizeNumbers(lhs, rhs)
	if err != nil {
		return err
	}

	if keepFloat {
		stack.Push(lhsf - rhsf)
	} else {
		stack.Push(int64(lhsf - rhsf))
	}

	return nil
}

func OPEqual(stack *core.Stack, _ core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	var bitSet int64

	// TODO: at some point this should be using bool type
	if lhs == rhs {
		bitSet = 1
	}

	stack.Push(bitSet)

	return nil
}

func OPDump(stack *core.Stack, _ core.Operation) error {
	value, err := stack.Pop()
	if err != nil {
		return err
	}

	// TODO: Remove the new line character
	fmt.Println(value)

	return nil
}
