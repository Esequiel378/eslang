package tokens

import (
	ops "eslang/core/operations"
)

// TokenDump function  Óòß  dumps the stack
func TokenDump(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DUMP)

	err := program.Push(op)

	return true, err
}
