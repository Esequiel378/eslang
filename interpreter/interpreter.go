package interpreter

import (
	"eslang/core"
	"eslang/interpreter/stack"
	"fmt"
)

func executeProgram(program *core.Program, stack *stack.Stack) error {
	for _, op := range program.Operations() {
		switch op.Type() {
		case core.OP_BLOCK:
			handler, ok := REGISTERED_OP_BLOCK[op.TokenStart().Token()]
			if !ok {
				return FormatError(
					op,
					fmt.Errorf("exaustive block operation handiling for `%s`", op.TypeAlias()),
				)
			}

			for {
				program, err := handler(stack, op)
				if err != nil {
					return FormatError(op, err)
				}

				if program == nil {
					break
				}

				if err := executeProgram(program, stack); err != nil {
					return FormatError(op, err)
				}

				if op.TokenStart().Token() != core.TOKEN_WHILE {
					break
				}
			}

		default:
			handler, ok := REGISTERED_OPERATIONS[op.Type()]

			if !ok {
				return FormatError(
					op,
					fmt.Errorf("exaustive operation handiling for `%s`", op.TypeAlias()),
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
