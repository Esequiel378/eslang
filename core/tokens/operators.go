package tokens

import ops "eslang/core/operations"

// ======================
// ARITHMETIC OPERATIONS
// ======================

// TokenOperatorAdd function    Push the addition of the two topmost values on the stack onto the stack
func TokenOperatorAdd(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "+" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_ADD)

	err := program.Push(op)

	return true, err
}

// TokenOperatorSub function    Push the subtraction of the two topmost values on the stack onto the stack
func TokenOperatorSub(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "-" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_SUB)

	err := program.Push(op)

	return true, err
}

// TokenOperatorMul function    Push the multiplication of the two topmost values on the stack onto the stack
func TokenOperatorMul(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "*" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_MUL)

	err := program.Push(op)

	return true, err
}

// TokenOperatorDiv function    Push the division of the two topmost values on the stack onto the stack
func TokenOperatorDiv(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "/" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_OPERATOR_DIV)

	err := program.Push(op)

	return true, err
}

// TokenOperatorMod function    Push the modulo of the two topmost values on the stack onto the stack
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

// TokenROperatorEqual function    Push the equal comparition between the two topmost values on the stack onto the stack
func TokenROperatorEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_EQUAL)

	err := program.Push(op)

	return true, err
}

// TokenROperatorNotEqual function    Push the not equal comparition between the two topmost values on the stack onto the stack
func TokenROperatorNotEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "!=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_NOT_EQUAL)

	err := program.Push(op)

	return true, err
}

// TokenROperatorLessThan function    Push the less than comparition between the two topmost values on the stack onto the stack
func TokenROperatorLessThan(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "<" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_LESS_THAN)

	err := program.Push(op)

	return true, err
}

// TokenROperatorLessThanOrEqual function    Push the less than or equal comparition between the two topmost values on the stack onto the stack
func TokenROperatorLessThanOrEqual(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "<=" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_LESS_THAN_OR_EQUAL)

	err := program.Push(op)

	return true, err
}

// TokenROperatorGreaterThan function    Push the greater than comparition between the two topmost values on the stack onto the stack
func TokenROperatorGreaterThan(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != ">" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_R_OPERATOR_GREATER_THAN)

	err := program.Push(op)

	return true, err
}

// TokenROperatorGreaterThanOrEqual function    Push the greater than or equal comparition between the two topmost values on the stack onto the stack
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

// TokenLOperatorOr function    Push the logical or between the two topmost values on the stack onto the stack
func TokenLOperatorOr(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "||" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_L_OPERATOR_OR)

	err := program.Push(op)

	return true, err
}

// TokenLOperatorAnd function    Push the logical and between the two topmost values on the stack onto the stack
func TokenLOperatorAnd(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "&&" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_L_OPERATOR_AND)

	err := program.Push(op)

	return true, err
}

// TokenLOperatorNot function    Push the logical not between the two topmost values on the stack onto the stack
func TokenLOperatorNot(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "!" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewMiscOperation(position, ops.OP_L_OPERATOR_NOT)

	err := program.Push(op)

	return true, err
}
