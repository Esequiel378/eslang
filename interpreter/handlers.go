package interpreter

import (
	"eslang/core"
	"fmt"
)

func OPPush(stack *core.Stack, op core.Operation) error {
	stack.Push(op.Value())

	return nil
}

func OPPlus(stack *core.Stack, op core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

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

func OPMinus(stack *core.Stack, op core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

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

func OPEqual(stack *core.Stack, op core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	var bitSet int8

	// TODO: at some point this should be using bool type
	if lhs == rhs {
		bitSet = 1
	}

	stack.Push(bitSet)

	return nil
}

func OPDump(stack *core.Stack, op core.Operation) error {
	value, err := stack.Pop()

	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}
