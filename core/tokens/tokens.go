package tokens

import (
	ops "eslang/core/operations"
)

type (
	// TokenType int    is the type of a token.
	TokenType int
	// TokenHandler func    function definition to implement new token handlers
	TokenHandler func(token string, lnum, cnum int, program *ops.Program) (bool, error)
)

const (
	// blocks
	TOKEN_BLOCK_IF_ELSE TokenType = iota
	TOKEN_BLOCK_END
	TOKEN_BLOCK_WHILE

	// std
	TOKEN_DUMP

	// push
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STR

	// variables
	TOKEN_SET_VARIABLE_TYPE
	TOKEN_VARIABLE
	TOKEN_VARIABLE_WRITE

	// operators
	TOKEN_OPERATOR_ADD
	TOKEN_OPERATOR_SUB
	TOKEN_OPERATOR_MUL
	TOKEN_OPERATOR_DIV
	TOKEN_OPERATOR_MOD

	TOKEN_TYPE_COUNT
)

// REGISTERED_TOKENS map    is a map of token types to their respective token handlers
var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	// blocks
	TOKEN_BLOCK_END:     TokenBlockEnd,
	TOKEN_BLOCK_WHILE:   TokenBlockWhile,
	TOKEN_BLOCK_IF_ELSE: TokenBlockIfElse,

	// std
	TOKEN_DUMP: TokenDump,

	// push
	TOKEN_PUSH_FLOAT: TokenPushFloat,
	TOKEN_PUSH_INT:   TokenPushInt,
	TOKEN_PUSH_STR:   TokenPushStr,

	// variables
	TOKEN_SET_VARIABLE_TYPE: TokenSetVariableType,
	TOKEN_VARIABLE:          TokenVariable,
	TOKEN_VARIABLE_WRITE:    TokenVariableWrite,

	// operators
	TOKEN_OPERATOR_ADD: TokenOperatorAdd,
	TOKEN_OPERATOR_SUB: TokenOperatorSub,
	TOKEN_OPERATOR_MUL: TokenOperatorMul,
	TOKEN_OPERATOR_DIV: TokenOperatorDiv,
	TOKEN_OPERATOR_MOD: TokenOperatorMod,
}
