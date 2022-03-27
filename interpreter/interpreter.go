package interpreter

import (
	ops "eslang/core/operations"
	"eslang/interpreter/handlers"
	"eslang/interpreter/stack"
	"fmt"
)

func executeProgram(program *ops.Program, stack *stack.Stack) error {
	for _, op := range program.Operations() {
		switch op.Type() {
		default:
			handler, ok := handlers.REGISTERED_OPERATIONS[op.Type()]

			if !ok {
				return FormatError(
					op,
					fmt.Errorf("exaustive operation handiling for `%s`", op.Type()),
				)
			}

			if err := handler(stack, op); err != nil {
				return FormatError(op, err)
			}
		}
	}

	return nil
}

func SimulateProgram(program *ops.Program) error {
	stack := stack.NewStack()

	if err := executeProgram(program, &stack); err != nil {
		return err
	}

	return nil
}
