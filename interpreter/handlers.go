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

	stack.Push(lhs.(int) + rhs.(int))

	return nil
}

func OPMinus(stack *core.Stack, op *core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	stack.Push(lhs.(int) - rhs.(int))

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
