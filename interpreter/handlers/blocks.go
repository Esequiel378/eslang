package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

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
