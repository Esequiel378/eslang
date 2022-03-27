package tokens

import (
	ops "eslang/core/operations"
)

type (
	// TokenType int    is the type of a token.
	TokenType int
	// TokenHandler func    function definition to implement new token handlers
	TokenHandler func(token string, lnum, column int, program *ops.Program) (bool, error)
)

const (
	TOKEN_DUMP TokenType = iota
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STR

	TOKEN_TYPE_COUNT
)

// REGISTERED_TOKENS map    is a map of token types to their respective token handlers
var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	TOKEN_DUMP:       TokenDump,
	TOKEN_PUSH_FLOAT: TokenPushFloat,
	TOKEN_PUSH_INT:   TokenPushInt,
	TOKEN_PUSH_STR:   TokenPushStr,
}

// TOKEN_ALIASES map    is a map of token types to their respective string representations
var TOKEN_ALIASES = map[TokenType]string{
	TOKEN_DUMP:       "DUMP",
	TOKEN_PUSH_FLOAT: "PUSH_FLOAT",
	TOKEN_PUSH_INT:   "PUSH_INT",
	TOKEN_PUSH_STR:   "PUSH_STR",
}

// String method    returns the string representation of a token type
func (t TokenType) String() string {
	return TOKEN_ALIASES[t]
}
