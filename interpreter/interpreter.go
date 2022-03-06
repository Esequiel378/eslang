package interpreter

import (
	"eslang/core"
	"fmt"
	"reflect"
)

func FormatError(op core.Operation, err error) error {
	line, col := op.TokenStart().Position()
	token := op.TokenStart().TokenAlias()

	return fmt.Errorf("error %s with Token: %s int line: %d/%d", err.Error(), token, line, col)
}

func executeOperations(program *core.Program, stack *core.Stack) error {
	for _, op := range *program {
		handler := REGISTERED_OPERATIONS[op.Type()]

		switch op.Type() {
		case core.OP_BLOCK:
			program, ok := op.Value().(*core.Program)

			if !ok {
				return FormatError(op, fmt.Errorf("error running block program"))
			}

			if err := executeOperations(program, stack); err != nil {
				return FormatError(op, err)
			}
		case core.OP_IF:
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

			if isNil(block.ElseBlock()) {
				break
			}

			program := block.ElseBlock()

			if err := executeOperations(program, stack); err != nil {
				return FormatError(op, err)
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
