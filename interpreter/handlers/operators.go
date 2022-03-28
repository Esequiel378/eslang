package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

// ======================
// ARITHMETIC OPERATIONS
// ======================

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

// ======================
// RELATIONAL OPERATIONS
// ======================

// OPREqual
// OPRNotEqual
// OPRLessThan
// OPRLessThanOrEqual
// OPRGreaterThan
// OPRGreaterThanOrEqual

// OPREqual function    compare if two values are equal and push 1 onto the stack if they are, 0 otherwise
func OPREqual(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.EqualValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPRNotEqual function    compare if two values are not equal and push 1 onto the stack if they are, 0 otherwise
func OPRNotEqual(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.NotEqualValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPRLessThan function    compare if the first value is less than the second and push 1 onto the stack if it is, 0 otherwise
func OPRLessThan(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.LessThanValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPRGreaterThan function    compare if the first value is greater than the second and push 1 onto the stack if it is, 0 otherwise
func OPRGreaterThan(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.GreaterThanValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPRLessThanOrEqual function    compare if the first value is less than or equal to the second and push 1 onto the stack if it is, 0 otherwise
func OPRLessThanOrEqual(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.LessThanOrEqualValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

// OPRGreaterThanOrEqual function    compare if the first value is greater than or equal to the second and push 1 onto the stack if it is, 0 otherwise
func OPRGreaterThanOrEqual(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	value, err := s.GreaterThanOrEqualValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}
