package interpreter

import (
	ops "eslang/core/operations"
	"eslang/interpreter/handlers"
	"eslang/interpreter/stack"
)

func executeProgram(program *ops.Program, stack *stack.Stack) error {
	for _, op := range program.Operations() {
		switch op.Type() {
		case ops.OP_BLOCK_IF_ELSE:
			handler := handlers.REGISTERED_BLOCK_OPERATIONS[op.Type()]

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

		case ops.OP_BLOCK_WHILE:
			handler := handlers.REGISTERED_BLOCK_OPERATIONS[ops.OP_BLOCK_WHILE]

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
			}

		default:
			handler := handlers.REGISTERED_OPERATIONS[op.Type()]

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
