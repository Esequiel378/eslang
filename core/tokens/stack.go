package tokens

import (
	ops "eslang/core/operations"
)

// TokenDrop function    drops the top of the stack
func TokenDrop(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "drop" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DROP)

	err := program.Push(op)

	return true, err
}

// TokenDump function    dumps the stack
func TokenDump(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DUMP)

	err := program.Push(op)

	return true, err
}

// TokenDup function    duplicates the top of the stack
func TokenDup(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "dup" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DUP)

	err := program.Push(op)

	return true, err
}

// TokenOver function    duplicate the second-to-top of the stack
func TokenOver(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "over" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OVER)

	err := program.Push(op)

	return true, err
}

// TokenSwap function    swaps the top two elements of the stack
func TokenSwap(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "swap" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_SWAP)

	err := program.Push(op)

	return true, err
}
