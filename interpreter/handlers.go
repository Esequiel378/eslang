package interpreter

import (
	"eslang/core"
	"fmt"
)

func OPPush(stack *core.Stack, op *core.Operation) error {
	stack.Push(op.Value)

	return nil
}

func OPPlus(stack *core.Stack, op *core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	lhsf, rhsf, err := normalizeNumbers(lhs, rhs)

	if err != nil {
		return err
	}

	keepFloat := false

	numbers := []interface{}{lhs, rhs}

	for _, number := range numbers {
		if _, ok := number.(float64); ok {
			keepFloat = true
			break
		}
	}

	if !keepFloat {
		stack.Push(int64(lhsf + rhsf))
	} else {
		stack.Push(lhsf + rhsf)
	}

	return nil
}

func OPMinus(stack *core.Stack, op *core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	lhsf, rhsf, err := normalizeNumbers(lhs, rhs)

	if err != nil {
		return err
	}

	keepFloat := false

	numbers := []interface{}{lhs, rhs}

	for _, number := range numbers {
		if _, ok := number.(float64); ok {
			keepFloat = true
			break
		}
	}

	if !keepFloat {
		stack.Push(int64(lhsf - rhsf))
	} else {
		stack.Push(lhsf - rhsf)
	}

	return nil
}

func OPDump(stack *core.Stack, op *core.Operation) error {
	value, err := stack.Pop()

	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}
