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
	TOKEN_BLOCK_IF_ELSE TokenType = iota
	TOKEN_BLOCK_END
	TOKEN_DUMP
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STR
	TOKEN_SET_VARIABLE_TYPE
	TOKEN_VARIABLE
	TOKEN_VARIABLE_WRITE

	TOKEN_TYPE_COUNT
)

// REGISTERED_TOKENS map    is a map of token types to their respective token handlers
var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	TOKEN_BLOCK_END:         TokenBlockEnd,
	TOKEN_BLOCK_IF_ELSE:     TokenBlockIfElse,
	TOKEN_DUMP:              TokenDump,
	TOKEN_PUSH_FLOAT:        TokenPushFloat,
	TOKEN_PUSH_INT:          TokenPushInt,
	TOKEN_PUSH_STR:          TokenPushStr,
	TOKEN_SET_VARIABLE_TYPE: TokenSetVariableType,
	TOKEN_VARIABLE:          TokenVariable,
	TOKEN_VARIABLE_WRITE:    TokenVariableWrite,
}
