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
	// std
	ops.OP_DUMP: OPDump,

	// push
	ops.OP_PUSH_FLOAT:  OPPushFloat,
	ops.OP_PUSH_INT:    OPPushInt,
	ops.OP_PUSH_STRING: OPPushStr,

	// variables
	ops.OP_VARIABLE:       OPVariable,
	ops.OP_VARIABLE_WRITE: OPVariableWrite,

	// operators
	ops.OP_OPERATOR_ADD: OPOperatorAdd,
	ops.OP_OPERATOR_SUB: OPOperatorSub,
	ops.OP_OPERATOR_MUL: OPOperatorMul,
	ops.OP_OPERATOR_DIV: OPOperatorDiv,
	ops.OP_OPERATOR_MOD: OPOperatorMod,
}

// REGISTERED_BLOCK_OPERATIONS    is a map of block operation handlers
var REGISTERED_BLOCK_OPERATIONS = map[ops.OPType]OPBlockHandler{
	ops.OP_BLOCK_IF_ELSE: OPBlockIfElse,
	ops.OP_BLOCK_WHILE:   OPBlockWhile,
}
