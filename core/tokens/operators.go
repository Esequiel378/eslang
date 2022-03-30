package tokens

import ops "eslang/core/operations"

// ======================
// ARITHMETIC OPERATIONS
// ======================

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

// ======================
// RELATIONAL OPERATIONS
// ======================

// TokenROperatorEqual function    push the equal operation onto the stack
func TokenROperatorEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_EQUAL)

	err := program.Push(op)

	return true, err
}

// TokenROperatorNotEqual function    push the not equal operation onto the stack
func TokenROperatorNotEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "!=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_NOT_EQUAL)

	err := program.Push(op)

	return true, err
}

// TokenROperatorLessThan function    push the less than operation onto the stack
func TokenROperatorLessThan(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "<" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_LESS_THAN)

	err := program.Push(op)

	return true, err
}

// TokenROperatorLessThanOrEqual function    push the less than or equal operation onto the stack
func TokenROperatorLessThanOrEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "<=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_LESS_THAN_OR_EQUAL)

	err := program.Push(op)

	return true, err
}

// TokenROperatorGreaterThan function    push the greater than operation onto the stack
func TokenROperatorGreaterThan(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != ">" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_GREATER_THAN)

	err := program.Push(op)

	return true, err
}

// TokenROperatorGreaterThanOrEqual function    push the greater than or equal operation onto the stack
func TokenROperatorGreaterThanOrEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != ">=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_GREATER_THAN_OR_EQUAL)

	err := program.Push(op)

	return true, err
}

// ==================
// LOGICAL OPERATIONS
// ==================

// TokenLOperatorOr function    push the or operation onto the stack
func TokenLOperatorOr(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "||" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_L_OPERATOR_OR)

	err := program.Push(op)

	return true, err
}

// TokenLOperatorAnd function    push the and operation onto the stack
func TokenLOperatorAnd(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "&&" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_L_OPERATOR_AND)

	err := program.Push(op)

	return true, err
}

// TokenLOperatorNot function    push the not operation onto the stack
func TokenLOperatorNot(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "!" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_L_OPERATOR_NOT)

	err := program.Push(op)

	return true, err
}
