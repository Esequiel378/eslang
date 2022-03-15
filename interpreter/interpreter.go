package interpreter

import (
	"eslang/core"
	"fmt"
	"reflect"
)

func FormatError(op core.Operation, err error) error {
	line, col := op.TokenStart().Position()
	token := op.TokenStart().TokenAlias()

	return fmt.Errorf("error `%s` with Token: %s in line: %d/%d", err.Error(), token, line, col)
}

func executeOperations(program *core.Program, stack *core.Stack) error {
	for _, op := range *program {
		handler := REGISTERED_OPERATIONS[op.Type()]

		switch op.Type() {
		case core.OP_BLOCK:
			switch op.TokenStart().Token() {
			case core.TOKEN_DO:
				program, ok := op.Value().(*core.Program)

				if !ok {
					return FormatError(op, fmt.Errorf("error running block program"))
				}

				if err := executeOperations(program, stack); err != nil {
					return FormatError(op, err)
				}
			case core.TOKEN_IF:
				value, err := stack.Pop()
				if err != nil {
					return FormatError(op, fmt.Errorf("error running block program"))
				}

				truthy, ok := value.(int64)

				if !ok {
					return FormatError(op, fmt.Errorf(
						"error testing the truthy of %s with type %s",
						value,
						reflect.TypeOf(value),
					),
					)
				}

				block := op.(core.BlockOperation)

				if truthy != 0 {
					program, ok := block.Value().(*core.Program)

					if !ok {
						return FormatError(op, fmt.Errorf("error running block program"))
					}

					if err := executeOperations(program, stack); err != nil {
						return FormatError(op, err)
					}

					break
				}

				if isNil(block.RefBlock()) {
					break
				}

				program := block.RefBlock().Block()

				if err := executeOperations(program, stack); err != nil {
					return FormatError(op, err)
				}

			}

		default:
			if err := handler(stack, op); err != nil {
				return FormatError(op, err)
			}
		}
	}

	return nil
}

func SimulateProgram(program *core.Program) error {
	var stack core.Stack

	if err := executeOperations(program, &stack); err != nil {
		return err
	}

	return nil
}
