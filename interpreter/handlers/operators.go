package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

// OPOperatorAdd function    add two values and push the result onto the stack
func OPOperatorAdd(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.AddValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPOperatorSub function    subtract two values and push the result onto the stack
func OPOperatorSub(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.SubtractValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPOperatorMul function    multiply two values and push the result onto the stack
func OPOperatorMul(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.MultiplyValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPOperatorDiv function    divide two values and push the result onto the stack
func OPOperatorDiv(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.DivideValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPOperatorMod function    modulo two values and push the result onto the stack
func OPOperatorMod(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.ModuloValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}
