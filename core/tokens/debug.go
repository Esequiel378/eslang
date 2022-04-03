package tokens

import ops "eslang/core/operations"

// TokenDebug function    Prints the stack
func TokenDebug(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "debug" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DEBUG)

	err := program.Push(op)

	return true, err
}
