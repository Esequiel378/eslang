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
		ifProgram := program.Copy()
		op := ops.NewOPBlockIfElse(ifProgram, position)
		err := program.Push(op)

		return true, err
	}

	if token == "else" {
		lastOp := program.LastOP()

		if lastOp.Type() != ops.OP_BLOCK_IF_ELSE {
			return false, fmt.Errorf("`else` must be used after `if` block")
		}

		ifBlock := lastOp.(*ops.OPBlockIfElse)

		ifBlock.CloseLastBlock()

		elseProgram := program.Copy()
		esleBlock := ops.NewOPBlockIfElse(elseProgram, position)

		ifBlock.SetNext(esleBlock)

		return true, nil
	}

	return false, nil
}

func TokenBlockWhile(token string, lnum, column int, program *ops.Program) (bool, error) {
	if token != "while" {
		return false, nil
	}

	position := ops.NewPosition(lnum, column, "")
	whileProgram := program.Copy()

	op := ops.NewOPBlockWhile(whileProgram, position)

	err := program.Push(op)

	return true, err
}

// TokenBlockEnd function    closes a block of code that can be executed
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
