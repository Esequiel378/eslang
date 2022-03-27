package tokens

import (
	ops "eslang/core/operations"
	"fmt"
)

// TokenBlockIfElse function    pushes a block of code that can be executed
// if a condition is met onto the stack
func TokenBlockIfElse(token string, lnum, column int, program *ops.Program) (bool, error) {
	position := ops.NewPosition(lnum, column, "")

	switch token {
	case "if":
		op := ops.NewOPBlockIfElse(position)
		program.Push(op)

	case "else":
		lastOp := program.LastOP()

		if lastOp.Type() != ops.OP_BLOCK_IF_ELSE {
			return false, fmt.Errorf("`else` must be used after `if` block")
		}

		ifBlock := lastOp.(*ops.OPBlockIfElse)

		ifBlock.CloseLastBlock()
		ifBlock.SetNext(position)

	default:
		return false, nil
	}

	return true, nil
}

func TokenBlockEnd(token string, _, _ int, program *ops.Program) (bool, error) {
	if token != "end" {
		return false, nil
	}

	lastOp := program.LastOP()

	if !lastOp.Type().IsBlock() {
		return false, fmt.Errorf("`end` must be used to close a block")
	}

	block := lastOp.(ops.OperationBlock)
	block.CloseBlock()

	return true, nil
}
