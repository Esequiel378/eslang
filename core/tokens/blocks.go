package tokens

import (
	ops "eslang/core/operations"
	"fmt"
)

// TokenBlockIfElse function    Push a block of code that can be executed if a condition is met onto the stack
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
		var lastBlock ops.OperationBlock

		if lastOp.Type().IsBlock() {
			lastBlock = lastOp.(ops.OperationBlock).LastNestedBlock()
		}

		if lastBlock.Type() != ops.OP_BLOCK_IF_ELSE {
			return false, fmt.Errorf("`else` must be used after `if` block")
		}

		ifBlock := lastBlock.(*ops.OPBlockIfElse)

		ifBlock.CloseLastBlock()

		elseProgram := program.Copy()
		esleBlock := ops.NewOPBlockIfElse(elseProgram, position)

		ifBlock.SetNext(esleBlock)

		return true, nil
	}

	return false, nil
}

// TokenBlockWhile function    Push a block of code that can be executed while a condition is met onto the stack
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
func TokenBlockEnd(token string, lnum, column int, program *ops.Program) (bool, error) {
	if token != "end" {
		return false, nil
	}

	lastOp := program.LastOP()

	if !lastOp.Type().IsBlock() {
		return true, fmt.Errorf("`end` must be used after a block operation in line %d:%d", lnum, column)
	}

	lastBlock := lastOp.(ops.OperationBlock).LastNestedBlock()

	lastBlock.CloseBlock()

	return true, nil
}

// TokenDo function    starts a block of code that can be executed closing a previous condition block
func TokenDo(token string, lnum, column int, program *ops.Program) (bool, error) {
	if token != "do" {
		return false, nil
	}

	lastOp := program.LastOP()

	if !lastOp.Type().IsBlock() {
		return true, fmt.Errorf("`do` must be used after a loop operation in line %d:%d", lnum, column)
	}

	lastBlock := lastOp.(ops.OperationBlock).LastNestedBlock()

	block, ok := lastBlock.(ops.OperationLoop)
	if !ok {
		return true, fmt.Errorf("`do` must be used after a loop operation")
	}

	block.CloseConditionBlock()

	return true, nil
}
