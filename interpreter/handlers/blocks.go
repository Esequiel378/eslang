package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

// OPBlockIfElse function    return the program to run for a if-else operation
func OPBlockIfElse(stack *s.Stack, op ops.Operation) (*ops.Program, error) {
	sValue, err := stack.Peek()
	if err != nil {
		return nil, err
	}

	truthy, err := sValue.TestTruthy()
	if err != nil {
		return nil, err
	}

	block := op.(*ops.OPBlockIfElse)

	// `if` block
	if truthy {
		return block.Program(), nil
	}

	// `else` block
	if block.HasNext() {
		next := block.Next().(*ops.OPBlockIfElse)
		program := next.Program()

		return program, nil
	}

	return nil, nil
}

// OPBlockWhile function    return the program to run for a while operation
func OPBlockWhile(stack *s.Stack, op ops.Operation) (*ops.Program, error) {
	sValue, err := stack.Peek()
	if err != nil {
		return nil, err
	}

	truthy, err := sValue.TestTruthy()
	if err != nil {
		return nil, err
	}

	block := op.(*ops.OPBlockWhile)

	// `while` block
	if truthy {
		return block.Program(), nil
	}

	return nil, nil
}
