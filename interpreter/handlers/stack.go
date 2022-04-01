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
	sValue, err := stack.Peek()
	if err != nil {
		return err
	}

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

// OPRot function    rotates the top three items on the stack
// TODO: make this operation more efficient
func OPRot(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	rotValue, err := stack.Pop()
	if err != nil {
		return err
	}

	stack.Push(lhs)
	stack.Push(rhs)

	stack.Push(rotValue)

	return nil
}

// OPORot function    rotates the top three items on the stack in the opposite direction
func OPORot(stack *s.Stack, _ ops.Operation) error {
	rotValue, err := stack.Pop()
	if err != nil {
		return err
	}

	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(rotValue)

	stack.Push(lhs)
	stack.Push(rhs)

	return nil
}

// OPTuck function    duplicates the top of the stack and places it below the second-to-top
func OPTuck(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(rhs)
	stack.Push(lhs)
	stack.Push(rhs)

	return nil
}

// OPTwoDup function    duplicates the top two items on the stack
func OPTwoDup(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PeekTwo()
	if err != nil {
		return err
	}

	stack.Push(lhs)
	stack.Push(rhs)

	return nil
}
