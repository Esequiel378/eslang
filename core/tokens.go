package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type (
	TokenType    int
	TokenHandler func(token, line string, lnum int, program *Program) error
)

const (
	TOKEN_DO TokenType = iota
	TOKEN_DUMP
	TOKEN_ELSE
	TOKEN_END
	TOKEN_EQUAL
	TOKEN_IF
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_PUSH
)

var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	TOKEN_DO:    TokenDo,
	TOKEN_DUMP:  TokenDump,
	TOKEN_ELSE:  TokenElse,
	TOKEN_END:   TokenEnd,
	TOKEN_EQUAL: TokenEqual,
	TOKEN_IF:    TokenIf,
	TOKEN_MINUS: TokenMinus,
	TOKEN_PLUS:  TokenPlus,
	TOKEN_PUSH:  TokenPush,
}

var TOKEN_MAPPING = map[TokenType]string{
	TOKEN_DO:    "DO",
	TOKEN_DUMP:  "DUMP",
	TOKEN_ELSE:  "ELSE",
	TOKEN_END:   "END",
	TOKEN_EQUAL: "EQUAL",
	TOKEN_IF:    "IF",
	TOKEN_MINUS: "MINUS",
	TOKEN_PLUS:  "PLUS",
	TOKEN_PUSH:  "PUSH",
}

var IS_DIGIT = regexp.MustCompile(`^[0-9]\d*(\.\d+)?$`)

type Token struct {
	line  int
	col   int
	token TokenType
}

func (t *Token) Token() TokenType {
	return t.token
}

func (t *Token) TokenAlias() string {
	return TOKEN_MAPPING[t.token]
}

func (t *Token) Position() (int, int) {
	return t.line, t.col
}

func (t *Token) SetPostition(line, col int) {
	t.line = line
	t.col = col
}

func tokenPushFloat64(token string) (*Operation, error) {
	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	opValue := NewOperationValue().SetFloat(value)
	op := NewOperation(OP_PUSH, opValue, TOKEN_PUSH, TOKEN_PUSH)

	return op, nil
}

func tokenPushInt64(token string) (*Operation, error) {
	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	opValue := NewOperationValue().SetInt(value)
	op := NewOperation(OP_PUSH, opValue, TOKEN_PUSH, TOKEN_PUSH)

	return op, nil
}

func TokenPush(token, line string, lnum int, program *Program) error {
	if !IS_DIGIT.MatchString(token) {
		return fmt.Errorf("invalid token")
	}

	var op *Operation
	var err error

	if strings.Contains(token, ".") {
		op, err = tokenPushFloat64(token)
	} else {
		op, err = tokenPushInt64(token)
	}

	if err != nil {
		return err
	}

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenPlus(token, line string, lnum int, program *Program) error {
	if token != "+" {
		return fmt.Errorf("invalid token")
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_MOP, opValue, TOKEN_PLUS, TOKEN_PLUS)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenMinus(token, line string, lnum int, program *Program) error {
	if token != "-" {
		return fmt.Errorf("invalid token")
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_MOP, opValue, TOKEN_MINUS, TOKEN_MINUS)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenEqual(token, line string, lnum int, program *Program) error {
	if token != "=" {
		return fmt.Errorf("invalid token")
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_MOP, opValue, TOKEN_EQUAL, TOKEN_EQUAL)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenDump(token, line string, lnum int, program *Program) error {
	if token != "." {
		return fmt.Errorf("invalid token")
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_DUMP, opValue, TOKEN_DUMP, TOKEN_DUMP)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenDo(token, line string, lnum int, program *Program) error {
	if token != "do" {
		return fmt.Errorf("invalid token")
	}

	cnum := strings.Index(line, token)

	opValue := NewOperationValue()
	opValue.Block().TokenStart().SetPostition(lnum+1, cnum+1)

	op := NewOperation(OP_BLOCK, opValue, TOKEN_DO, TOKEN_END)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenIf(token, line string, lnum int, program *Program) error {
	if token != "if" {
		return fmt.Errorf("invalid token")
	}

	cnum := strings.Index(line, token)

	opValue := NewOperationValue()
	opValue.Block().TokenStart().SetPostition(lnum+1, cnum+1)

	op := NewOperation(OP_BLOCK, opValue, TOKEN_IF, TOKEN_END)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return nil
}

func TokenElse(token, line string, lnum int, program *Program) error {
	if token != "else" {
		return fmt.Errorf("invalid token")
	}

	cnum := strings.Index(line, token)

	if err := program.CloseLastBlock(lnum+1, cnum+1); err != nil {
		return err
	}

	block := NewBlock(TOKEN_ELSE, TOKEN_END)
	block.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Last().Value().Block().SetNext(block)

	return nil
}

func TokenEnd(token, line string, lnum int, program *Program) error {
	if token != "end" {
		return fmt.Errorf("invalid token")
	}

	cnum := strings.Index(line, token)

	if err := program.CloseLastBlock(lnum+1, cnum+1); err != nil {
		return err
	}

	return nil
}
