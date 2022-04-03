package tokens

import (
	ops "eslang/core/operations"
)

// TokenDrop function    Drops the top of the stack
func TokenDrop(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "drop" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DROP)

	err := program.Push(op)

	return true, err
}

// TokenDump function    Dumps the stack
func TokenDump(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DUMP)

	err := program.Push(op)

	return true, err
}

// TokenDup function    Duplicates the top of the stack
func TokenDup(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "dup" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_DUP)

	err := program.Push(op)

	return true, err
}

// TokenOver function    Duplicate the second-to-top of the stack
func TokenOver(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "over" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OVER)

	err := program.Push(op)

	return true, err
}

// TokenSwap function    Swaps the top two elements of the stack
func TokenSwap(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "swap" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_SWAP)

	err := program.Push(op)

	return true, err
}

// TokenRot function    Rotates the top three elements of the stack
func TokenRot(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "rot" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_ROT)

	err := program.Push(op)

	return true, err
}

// TokenORot function    Rotates the top three elements of the stack in the opposite direction
func TokenORot(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "-rot" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_O_ROT)

	err := program.Push(op)

	return true, err
}

// TokenTuck function    Duplicates the top of the stack and places it below the second-to-top
func TokenTuck(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "tuck" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_TUCK)

	err := program.Push(op)

	return true, err
}

// TokenTwoDup function    Duplicates the top two elements of the stack
func TokenTwoDup(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "2dup" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_TWO_DUP)

	err := program.Push(op)

	return true, err
}

// TokenNip function    Drops the second-to-top element of the stack
func TokenNip(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "nip" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_NIP)

	err := program.Push(op)

	return true, err
}

// TokenTwoDrop function    Drops the top two elements of the stack
func TokenTwoDrop(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "2drop" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_TWO_DROP)

	err := program.Push(op)

	return true, err
}

// TokenTwoOver function    Duplicates the second-to-top two elements of the stack
func TokenTwoOver(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "2over" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_TWO_OVER)

	err := program.Push(op)

	return true, err
}

// TokenTwoSwap function    Swaps the second-to-top two elements of the stack
func TokenTwoSwap(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "2swap" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_TWO_SWAP)

	err := program.Push(op)

	return true, err
}
