package interpreter

import (
	"eslang/core"
	"fmt"
)

func FormatError(op core.Operation, err error) error {
	line, col := op.TokenStart().Position()
	token := op.TokenStart().Token()

	return fmt.Errorf("error %s with Token: %d int line: %d/%d", err.Error(), token, line, col)
}

func executeOperations(program *core.Program, stack *core.Stack) error {
	for _, op := range *program {
		handler := REGISTERED_OPERATIONS[op.Type()]

		if op.Type() == core.OP_BLOCK {
			program, err := op.Value(stack)

			if err != nil {
				return FormatError(op, err)
			}

			if program == nil {
				continue
			}

			pp, ok := program.(*core.Program)

			if !ok {
				return FormatError(op, fmt.Errorf("error running block program"))
			}

			if err := executeOperations(pp, stack); err != nil {
				return FormatError(op, err)
			}

			continue
		}

		if err := handler(stack, op); err != nil {
			return FormatError(op, err)
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
