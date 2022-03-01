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
	TOKEN_IF    = iota
	TOKEN_ELSE  = iota
	TOKEN_END   = iota
)

var REGISTERED_TOKENS = map[int]func(string, *BlockStack) (Operation, error){
	TOKEN_PUSH:  TokenPush,
	TOKEN_PLUS:  TokenPlus,
	TOKEN_MINUS: TokenMinus,
	TOKEN_EQUAL: TokenEqual,
	TOKEN_DUMP:  TokenDump,
	TOKEN_DO:    TokenDo,
	TOKEN_IF:    TokenIf,
	TOKEN_ELSE:  TokenElse,
	TOKEN_END:   TokenEnd,
}

var IS_DIGIT = regexp.MustCompile(`^[0-9]\d*(\.\d+)?$`)

func tokenPushFloat64(token string) (Operation, error) {
	value, err := strconv.ParseFloat(token, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	operation := NewOP(OP_PUSH, value)

	return operation, nil
}

func tokenPushInt64(token string) (Operation, error) {
	value, err := strconv.ParseInt(token, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	operation := NewOP(OP_PUSH, value)

	return operation, nil
}

func TokenPush(token string, blocks *BlockStack) (Operation, error) {
	if !IS_DIGIT.MatchString(token) {
		return nil, fmt.Errorf("Invalid token")
	}

	if strings.Contains(token, ".") {
		return tokenPushFloat64(token)
	}

	return tokenPushInt64(token)
}

func TokenPlus(token string, blocks *BlockStack) (Operation, error) {
	if token != "+" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := NewOP(OP_PLUS, nil)

	return operation, nil
}

func TokenMinus(token string, blocks *BlockStack) (Operation, error) {
	if token != "-" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := NewOP(OP_MINUS, nil)

	return operation, nil
}

func TokenEqual(token string, blocks *BlockStack) (Operation, error) {
	if token != "=" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := NewOP(OP_EQUAL, nil)

	return operation, nil
}

func TokenDump(token string, blocks *BlockStack) (Operation, error) {
	if token != "." {
		return nil, fmt.Errorf("Invalid token")
	}
	operation := NewOP(OP_DUMP, nil)

	return operation, nil
}

func TokenDo(token string, blocks *BlockStack) (Operation, error) {
	if token != "do" {
		return nil, fmt.Errorf("Invalid token")
	}

	blockOperation := NewMiscBlockOperation()

	blocks.Push(&blockOperation)

	return nil, nil
}

func TokenIf(token string, blocks *BlockStack) (Operation, error) {
	if token != "if" {
		return nil, fmt.Errorf("Invalid token")
	}

	ifBlockOperation := NewIfBlockOperation()

	blocks.Push(&ifBlockOperation)

	return nil, nil
}

func TokenElse(token string, blocks *BlockStack) (Operation, error) {
	if token != "else" {
		return nil, fmt.Errorf("Invalid token")
	}

	blocks.Last().EnableElseBlock()

	return nil, nil
}

func TokenEnd(token string, blocks *BlockStack) (Operation, error) {
	if token != "end" {
		return nil, fmt.Errorf("Invalid token")
	}

	block, err := blocks.Pop()

	if err != nil {
		return nil, err
	}

	return block, nil
}
