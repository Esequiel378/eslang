package core

import (
	ops "eslang/core/operations"
	"fmt"
	"regexp"
	"strconv"
)

type (
	TokenType int
	// TokenHandler type    function definition to implement new token handlers
	TokenHandler func(token string, lnum, column int, program *ops.Program) (bool, error)
)

const (
	TOKEN_DUMP TokenType = iota
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STR

	TOKEN_TYPE_COUNT
)

var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	TOKEN_DUMP:       TokenDump,
	TOKEN_PUSH_FLOAT: TokenPushFloat,
	TOKEN_PUSH_INT:   TokenPushInt,
	TOKEN_PUSH_STR:   TokenPushStr,
}

var TOKEN_ALIASES = map[TokenType]string{
	TOKEN_DUMP:       "DUMP",
	TOKEN_PUSH_FLOAT: "PUSH_FLOAT",
	TOKEN_PUSH_INT:   "PUSH_INT",
	TOKEN_PUSH_STR:   "PUSH_STR",
}

func (t TokenType) String() string {
	return TOKEN_ALIASES[t]
}

var (
	IS_INT   = regexp.MustCompile(`^\d+$`)
	IS_FLOAT = regexp.MustCompile(`^\d+\.\d+$`)
	IS_STR   = regexp.MustCompile(`^".+"$`)
)

func TokenPushInt(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_INT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushInt(value, position)

	program.Push(op)

	return true, nil
}

func TokenPushFloat(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_FLOAT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushFloat(value, position)

	program.Push(op)

	return true, nil
}

func TokenPushStr(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_STR.MatchString(token) {
		return false, nil
	}

	value := token[1 : len(token)-1]
	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushString(value, position)

	program.Push(op)

	return true, nil
}

func TokenDump(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPDump(position)

	program.Push(op)

	return true, nil
}
