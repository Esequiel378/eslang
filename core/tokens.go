package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	TOKEN_PUSH  = iota
	TOKEN_PLUS  = iota
	TOKEN_MINUS = iota
	TOKEN_EQUAL = iota
	TOKEN_DUMP  = iota
	TOKEN_DO    = iota
	TOKEN_END   = iota
)

var REGISTERED_TOKENS = map[int]func(string) (*Operation, error){
	TOKEN_PUSH:  TokenPush,
	TOKEN_PLUS:  TokenPlus,
	TOKEN_MINUS: TokenMinus,
	TOKEN_EQUAL: TokenEqual,
	TOKEN_DUMP:  TokenDump,
	TOKEN_DO:    TokenDo,
	TOKEN_END:   TokenEnd,
}

var IS_DIGIT = regexp.MustCompile(`^[1-9]\d*(\.\d+)?$`)

func tokenPushFloat64(token string) (*Operation, error) {
	value, err := strconv.ParseFloat(token, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	operation := Operation{
		Type:  OP_PUSH,
		Value: value,
	}

	return &operation, nil
}

func tokenPushInt64(token string) (*Operation, error) {
	value, err := strconv.ParseInt(token, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	operation := Operation{
		Type:  OP_PUSH,
		Value: value,
	}

	return &operation, nil
}

func TokenPush(token string) (*Operation, error) {
	if !IS_DIGIT.MatchString(token) {
		return nil, fmt.Errorf("Invalid token")
	}

	if strings.Contains(token, ".") {
		return tokenPushFloat64(token)
	}

	return tokenPushInt64(token)
}

func TokenPlus(token string) (*Operation, error) {
	if token != "+" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_PLUS,
	}

	return &operation, nil
}

func TokenMinus(token string) (*Operation, error) {
	if token != "-" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_MINUS,
	}

	return &operation, nil
}

func TokenEqual(token string) (*Operation, error) {
	if token != "=" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_EQUAL,
	}

	return &operation, nil
}

func TokenDump(token string) (*Operation, error) {
	if token != "." {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_DUMP,
	}

	return &operation, nil
}

func TokenDo(token string) (*Operation, error) {
	if token != "do" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_BLOCK,
	}

	return &operation, nil
}

func TokenEnd(token string) (*Operation, error) {
	if token != "end" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_BLOCK,
	}

	return &operation, nil
}
