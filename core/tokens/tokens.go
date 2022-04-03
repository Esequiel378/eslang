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
	TOKEN_DO

	// stack
	TOKEN_DROP
	TOKEN_DUMP
	TOKEN_DUP
	TOKEN_OVER
	TOKEN_O_ROT
	TOKEN_ROT
	TOKEN_SWAP
	TOKEN_TUCK
	TOKEN_TWO_DUP
	TOKEN_NIP
	TOKEN_TWO_DROP
	TOKEN_TWO_OVER
	TOKEN_TWO_SWAP

	// push
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STRING
	TOKEN_PUSH_BOOL

	// variables
	TOKEN_SET_VARIABLE_TYPE
	TOKEN_VARIABLE
	TOKEN_VARIABLE_WRITE

	// arithmetic operators
	TOKEN_OPERATOR_ADD
	TOKEN_OPERATOR_SUB
	TOKEN_OPERATOR_MUL
	TOKEN_OPERATOR_DIV
	TOKEN_OPERATOR_MOD

	// Relational operators
	TOKEN_R_OPERTATOR_EQUAL
	TOKEN_R_OPERTATOR_NOT_EQUAL
	TOKEN_R_OPERTATOR_LESS_THAN
	TOKEN_R_OPERTATOR_LESS_THAN_OR_EQUAL
	TOKEN_R_OPERTATOR_GREATER_THAN
	TOKEN_R_OPERTATOR_GREATER_THAN_OR_EQUAL

	// logical operators
	TOKEN_L_OPERTATOR_AND
	TOKEN_L_OPERTATOR_NOT
	TOKEN_L_OPERTATOR_OR

	TOKEN_TYPE_COUNT
)

// REGISTERED_TOKENS map    is a map of token types to their respective token handlers
var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	// blocks
	TOKEN_BLOCK_END:     TokenBlockEnd,
	TOKEN_BLOCK_WHILE:   TokenBlockWhile,
	TOKEN_BLOCK_IF_ELSE: TokenBlockIfElse,
	TOKEN_DO:            TokenDo,

	// stack
	TOKEN_DROP:     TokenDrop,
	TOKEN_DUMP:     TokenDump,
	TOKEN_DUP:      TokenDup,
	TOKEN_NIP:      TokenNip,
	TOKEN_OVER:     TokenOver,
	TOKEN_O_ROT:    TokenORot,
	TOKEN_ROT:      TokenRot,
	TOKEN_SWAP:     TokenSwap,
	TOKEN_TUCK:     TokenTuck,
	TOKEN_TWO_DROP: TokenTwoDrop,
	TOKEN_TWO_DUP:  TokenTwoDup,
	TOKEN_TWO_OVER: TokenTwoOver,
	TOKEN_TWO_SWAP: TokenTwoSwap,

	// push
	TOKEN_PUSH_FLOAT:  TokenPushFloat,
	TOKEN_PUSH_INT:    TokenPushInt,
	TOKEN_PUSH_STRING: TokenPushStr,
	TOKEN_PUSH_BOOL:   TokenPushBool,

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

	// relational opertators
	TOKEN_R_OPERTATOR_EQUAL:                 TokenROperatorEqual,
	TOKEN_R_OPERTATOR_NOT_EQUAL:             TokenROperatorNotEqual,
	TOKEN_R_OPERTATOR_LESS_THAN:             TokenROperatorLessThan,
	TOKEN_R_OPERTATOR_LESS_THAN_OR_EQUAL:    TokenROperatorLessThanOrEqual,
	TOKEN_R_OPERTATOR_GREATER_THAN:          TokenROperatorGreaterThan,
	TOKEN_R_OPERTATOR_GREATER_THAN_OR_EQUAL: TokenROperatorGreaterThanOrEqual,

	// logical operators
	TOKEN_L_OPERTATOR_AND: TokenLOperatorAnd,
	TOKEN_L_OPERTATOR_NOT: TokenLOperatorNot,
	TOKEN_L_OPERTATOR_OR:  TokenLOperatorOr,
}
