package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

type (
	// OPHandler    is a handler for operations
	OPHandler func(stack *s.Stack, _op ops.Operation) error

	// OPBinaryHandler    is a handler for block operation
	OPBlockHandler func(stack *s.Stack, _op ops.Operation) (*ops.Program, error)
)

// REGISTERED_OPERATIONS    is a map of operation handlers
var REGISTERED_OPERATIONS = map[ops.OPType]OPHandler{
	// stack
	ops.OP_DROP:     OPDrop,
	ops.OP_DUMP:     OPDump,
	ops.OP_DUP:      OPDup,
	ops.OP_NIP:      OPNip,
	ops.OP_OVER:     OPOver,
	ops.OP_O_ROT:    OPORot,
	ops.OP_ROT:      OPRot,
	ops.OP_SWAP:     OPSwap,
	ops.OP_TUCK:     OPTuck,
	ops.OP_TWO_DROP: OPTwoDrop,
	ops.OP_TWO_DUP:  OPTwoDup,
	ops.OP_TWO_OVER: OPTwoOver,
	ops.OP_TWO_SWAP: OPTwoSwap,

	// push
	ops.OP_PUSH_FLOAT:  OPPushFloat,
	ops.OP_PUSH_INT:    OPPushInt,
	ops.OP_PUSH_STRING: OPPushStr,
	ops.OP_PUSH_BOOL:   OPPushBool,

	// variables
	ops.OP_VARIABLE:       OPVariable,
	ops.OP_VARIABLE_WRITE: OPVariableWrite,

	// arithmetic operators
	ops.OP_OPERATOR_ADD: OPOperatorAdd,
	ops.OP_OPERATOR_SUB: OPOperatorSub,
	ops.OP_OPERATOR_MUL: OPOperatorMul,
	ops.OP_OPERATOR_DIV: OPOperatorDiv,
	ops.OP_OPERATOR_MOD: OPOperatorMod,

	// relational operators
	ops.OP_R_OPERATOR_EQUAL:                 OPREqual,
	ops.OP_R_OPERATOR_NOT_EQUAL:             OPRNotEqual,
	ops.OP_R_OPERATOR_LESS_THAN:             OPRLessThan,
	ops.OP_R_OPERATOR_LESS_THAN_OR_EQUAL:    OPRLessThanOrEqual,
	ops.OP_R_OPERATOR_GREATER_THAN:          OPRGreaterThan,
	ops.OP_R_OPERATOR_GREATER_THAN_OR_EQUAL: OPRGreaterThanOrEqual,

	// logical operators
	ops.OP_L_OPERATOR_AND: OPLOperatorAnd,
	ops.OP_L_OPERATOR_NOT: OPLOperatorNot,
	ops.OP_L_OPERATOR_OR:  OPLOperatorOr,

	// debug
	ops.OP_DEBUG: OPDebug,
}

// REGISTERED_BLOCK_OPERATIONS    is a map of block operation handlers
var REGISTERED_BLOCK_OPERATIONS = map[ops.OPType]OPBlockHandler{
	ops.OP_BLOCK_IF_ELSE: OPBlockIfElse,
	ops.OP_BLOCK_WHILE:   OPBlockWhile,
}
