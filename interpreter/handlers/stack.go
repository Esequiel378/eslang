package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

// OPDrop function    drop the top of the stack
func OPDrop(stack *s.Stack, _ ops.Operation) error {
	_, err := stack.Pop()

	return err
}

// OPDump function    dump the top of the stack
func OPDump(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	fmt.Println(sValue.Value())

	return nil
}

// OPDup function    duplicate the top of the stack
func OPDup(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	stack.Push(sValue)
	stack.Push(sValue)

	return nil
}

// OPOver function    duplicate the second-to-top of the stack
func OPOver(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.PeekAt(-2)
	if err != nil {
		return err
	}

	stack.Push(sValue)

	return nil
}

// OPSwap function    swap the top two items on the stack
func OPSwap(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(rhs)
	stack.Push(lhs)

	return nil
}