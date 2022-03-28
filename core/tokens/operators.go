package tokens

import ops "eslang/core/operations"

// TokenOperatorAdd function    push the add operation onto the stack
func TokenOperatorAdd(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "+" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_ADD)

	err := program.Push(op)

	return true, err
}

// TokenOperatorSub function    push the sub operation onto the stack
func TokenOperatorSub(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "-" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_SUB)

	err := program.Push(op)

	return true, err
}

// TokenOperatorMul function    push the mul operation onto the stack
func TokenOperatorMul(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "*" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_MUL)

	err := program.Push(op)

	return true, err
}

// TokenOperatorDiv function    push the div operation onto the stack
func TokenOperatorDiv(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "/" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_DIV)

	err := program.Push(op)

	return true, err
}

// TokenOperatorMod function    push the mod operation onto the stack
func TokenOperatorMod(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "%" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_MOD)

	err := program.Push(op)

	return true, err
}
