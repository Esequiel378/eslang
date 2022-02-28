package interpreter

import (
	"eslang/core"
)

var REGISTERED_OPERATIONS = map[int]func(*core.Stack, *core.Operation) error{
	core.OP_PUSH:  OPPush,
	core.OP_PLUS:  OPPlus,
	core.OP_MINUS: OPMinus,
	core.OP_DUMP:  OPDump,
}

func SimulateProgram(program *core.Program) error {
	var stack core.Stack

	for _, op := range *program {
		handler := REGISTERED_OPERATIONS[op.Type]

		if err := handler(&stack, op); err != nil {
			return err
		}
	}

	return nil
}
