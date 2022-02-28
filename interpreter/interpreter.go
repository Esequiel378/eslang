package interpreter

import (
	"eslang/core"
)

var REGISTERED_OPERATIONS = map[int]func(*core.Stack, *core.Operation) error{
	core.OP_PUSH:  OPPush,
	core.OP_PLUS:  OPPlus,
	core.OP_MINUS: OPMinus,
	core.OP_EQUAL: OPEqual,
	core.OP_DUMP:  OPDump,
}

func executeOperations(program *core.Program, stack *core.Stack) error {
	for _, op := range *program {
		handler := REGISTERED_OPERATIONS[op.Type]

		if op.Type == core.OP_BLOCK {
			if err := executeOperations(op.Value.(*core.Program), stack); err != nil {
				return err
			}

			continue
		}

		if err := handler(stack, op); err != nil {
			return err
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
