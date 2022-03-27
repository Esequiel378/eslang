package tokens

import (
	ops "eslang/core/operations"
	"fmt"
)

// TokenBlockIfElse function    pushes a block of code that can be executed
// if a condition is met onto the stack
func TokenBlockIfElse(token string, lnum, column int, program *ops.Program) (bool, error) {
	position := ops.NewPosition(lnum, column, "")

	if token == "if" {
		op := ops.NewOPBlockIfElse(position)

		program.Push(op)

		return true, nil
	}

	if token == "else" {
		lastOp := program.LastOP()

		if lastOp.Type() != ops.OP_BLOCK_IF_ELSE {
			return false, fmt.Errorf("`else` must be used after `if` block")
		}

		ifBlock := lastOp.(*ops.OPBlockIfElse)

		ifBlock.CloseLastBlock()
		ifBlock.SetNext(position)

		return true, nil
	}

	return false, nil
}
