package interpreter

import (
	"eslang/core"
	"eslang/interpreter/stack"
	"fmt"
)

func executeProgram(program *core.Program, stack *stack.Stack) error {
	for _, op := range program.Operations() {
		switch op.Type() {
		default:
			handler, ok := REGISTERED_OPERATIONS[op.Type()]

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

func SimulateProgram(program *core.Program) error {
	stack := stack.NewStack()

	if err := executeProgram(program, &stack); err != nil {
		return err
	}

	return nil
}
