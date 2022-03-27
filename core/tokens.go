package core

import (
	"fmt"
	"regexp"
	"strconv"
)

type (
	TokenType int
	// TokenHandler type    function definition to implement new token handlers
	TokenHandler func(token string, lnum, column int, program *Program) (bool, error)
)

const (
	TOKEN_DO TokenType = iota
	TOKEN_DUMP
	TOKEN_DUP
	TOKEN_ELSE
	TOKEN_END
	TOKEN_EQUAL
	TOKEN_FLOAT
	TOKEN_IF
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STR
	TOKEN_VAR
	TOKEN_VAR_READ
	TOKEN_VAR_TYPE
	TOKEN_VAR_WRITE
	TOKEN_WHILE

	TOKEN_TYPE_COUNT
)

var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	TOKEN_DUMP:       TokenDump,
	TOKEN_PUSH_FLOAT: TokenPushFloat,
	TOKEN_PUSH_INT:   TokenPushInt,
	TOKEN_PUSH_STR:   TokenPushStr,
}

var TOKEN_ALIASES = map[TokenType]string{
	TOKEN_DO:         "DO",
	TOKEN_DUMP:       "DUMP",
	TOKEN_DUP:        "DUP",
	TOKEN_ELSE:       "ELSE",
	TOKEN_END:        "END",
	TOKEN_EQUAL:      "EQUAL",
	TOKEN_IF:         "IF",
	TOKEN_MINUS:      "MINUS",
	TOKEN_PLUS:       "PLUS",
	TOKEN_PUSH_FLOAT: "PUSH_FLOAT",
	TOKEN_PUSH_INT:   "PUSH_INT",
	TOKEN_PUSH_STR:   "PUSH_STR",
	TOKEN_VAR:        "VAR",
	TOKEN_VAR_READ:   "VAR_READ",
	TOKEN_VAR_WRITE:  "VAR_WRITE",
	TOKEN_WHILE:      "WHILE",
}

func (t TokenType) String() string {
	if int(TOKEN_TYPE_COUNT)-1 != len(TOKEN_ALIASES) {
		panic("TOKEN_TYPE_COUNT exaust handling")
	}

	if alias, ok := TOKEN_ALIASES[t]; ok {
		return alias
	}

	return "-unknown-"
}

var (
	IS_INT            = regexp.MustCompile(`^\d+$`)
	IS_FLOAT          = regexp.MustCompile(`^\d+\.\d+$`)
	IS_STR            = regexp.MustCompile(`^".+"$`)
	IS_VALID_VARIABLE = regexp.MustCompile(`^[a-zA-Z]+$`)
)

func TokenPushInt(token string, lnum, cnum int, program *Program) (bool, error) {
	if !IS_INT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	position := NewPosition(lnum, cnum, "")
	op := NewOperationInt(value, position)

	program.Push(op)

	return true, nil
}

func TokenPushFloat(token string, lnum, cnum int, program *Program) (bool, error) {
	if !IS_FLOAT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	position := NewPosition(lnum, cnum, "")
	op := NewOperationFloat(value, position)

	program.Push(op)

	return true, nil
}

func TokenPushStr(token string, lnum, cnum int, program *Program) (bool, error) {
	if !IS_STR.MatchString(token) {
		return false, nil
	}

	value := token[1 : len(token)-1]
	position := NewPosition(lnum, cnum, "")
	op := NewOperationString(value, position)

	program.Push(op)

	return true, nil
}

func TokenDump(token string, lnum, cnum int, program *Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	position := NewPosition(lnum, cnum, "")
	op := NewOperationDump(position)

	program.Push(op)

	return true, nil
}
