package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

type (
	OPHandler func(*s.Stack, ops.Operation) error
)

var REGISTERED_OPERATIONS = map[ops.OPType]OPHandler{
	ops.OP_DUMP:        OPDump,
	ops.OP_PUSH_FLOAT:  OPPushFloat,
	ops.OP_PUSH_INT:    OPPushInt,
	ops.OP_PUSH_STRING: OPPushStr,
}
