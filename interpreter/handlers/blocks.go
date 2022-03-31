package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

// ExecuteBlockCondition function    execute a the condition of a block
func ExecuteBlockCondition(stack *s.Stack, operations []ops.Operation) (bool, error) {
	for _, op := range operations {
		if op.Type().IsBlock() {
			return false, fmt.Errorf("block operation not supported inside conditions")
		}

		// TODO: Maybe use a separated map for the operations that are allowed inside conditions
		handler := REGISTERED_OPERATIONS[op.Type()]

		if err := handler(stack, op); err != nil {
			return false, fmt.Errorf("error executing condition: %s", err)
		}
	}

	sValue, err := stack.Pop()
	if err != nil {
		return false, err
	}

	truthy, err := sValue.TestTruthy()

	return truthy, err
}

// OPBlockIfElse function    return the program to run for a if-else operation
func OPBlockIfElse(stack *s.Stack, op ops.Operation) (*ops.Program, error) {
	sValue, err := stack.Pop()
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
	block := op.(*ops.OPBlockWhile)

	truthy, err := ExecuteBlockCondition(stack, block.ConditionBlock())
	if err != nil {
		return nil, err
	}

	// `while` block
	if truthy {
		return block.Program(), nil
	}

	return nil, nil
}
