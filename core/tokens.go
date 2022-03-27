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
	op := NewOPPushInt(value, position)

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
	op := NewOPPushFloat(value, position)

	program.Push(op)

	return true, nil
}

func TokenPushStr(token string, lnum, cnum int, program *Program) (bool, error) {
	if !IS_STR.MatchString(token) {
		return false, nil
	}

	value := token[1 : len(token)-1]
	position := NewPosition(lnum, cnum, "")
	op := NewOPPushString(value, position)

	program.Push(op)

	return true, nil
}

func TokenDump(token string, lnum, cnum int, program *Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	position := NewPosition(lnum, cnum, "")
	op := NewOPDump(position)

	program.Push(op)

	return true, nil
}
