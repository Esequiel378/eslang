package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type TokenType int

const (
	TOKEN_DO    TokenType = iota
	TOKEN_DUMP            = iota
	TOKEN_ELSE            = iota
	TOKEN_END             = iota
	TOKEN_EQUAL           = iota
	TOKEN_IF              = iota
	TOKEN_MINUS           = iota
	TOKEN_PLUS            = iota
	TOKEN_PUSH            = iota
)

var REGISTERED_TOKENS = map[TokenType]func(token, line string, lnum int, blocks *BlockStack) (Operation, error){
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
	line  *int
	col   *int
	token TokenType
}

func (t *Token) Token() TokenType {
	return t.token
}

func (t *Token) TokenAlias() string {
	return TOKEN_MAPPING[t.token]
}

func (t *Token) Position() (int, int) {
	return *t.line, *t.col
}

func (t *Token) SetPostition(line, col int) {
	t.line = &line
	t.col = &col
}

func tokenPushFloat64(token string) (Operation, error) {
	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	operation := NewMiscOperation(OP_PUSH, value, TOKEN_PUSH)

	return operation, nil
}

func tokenPushInt64(token string) (Operation, error) {
	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	operation := NewMiscOperation(OP_PUSH, value, TOKEN_PUSH)

	return operation, nil
}

func TokenPush(token, line string, lnum int, _ *BlockStack) (op Operation, err error) {
	if !IS_DIGIT.MatchString(token) {
		return nil, fmt.Errorf("invalid token")
	}

	if strings.Contains(token, ".") {
		op, err = tokenPushFloat64(token)
	} else {
		op, err = tokenPushInt64(token)
	}

	cnum := strings.Index(line, token)

	op.TokenStart().SetPostition(lnum+1, cnum+1)

	return op, err
}

func TokenPlus(token, line string, lnum int, _ *BlockStack) (Operation, error) {
	if token != "+" {
		return nil, fmt.Errorf("invalid token")
	}

	op := NewMiscOperation(OP_MOP, nil, TOKEN_PLUS)
	cnum := strings.Index(line, token)

	op.TokenStart().SetPostition(lnum+1, cnum+1)

	return op, nil
}

func TokenMinus(token, line string, lnum int, _ *BlockStack) (Operation, error) {
	if token != "-" {
		return nil, fmt.Errorf("invalid token")
	}

	op := NewMiscOperation(OP_MOP, nil, TOKEN_MINUS)
	cnum := strings.Index(line, token)

	op.TokenStart().SetPostition(lnum+1, cnum+1)

	return op, nil
}

func TokenEqual(token, line string, lnum int, _ *BlockStack) (Operation, error) {
	if token != "=" {
		return nil, fmt.Errorf("invalid token")
	}

	op := NewMiscOperation(OP_MOP, nil, TOKEN_EQUAL)
	cnum := strings.Index(line, token)

	op.TokenStart().SetPostition(lnum+1, cnum+1)

	return op, nil
}

func TokenDump(token, line string, lnum int, _ *BlockStack) (Operation, error) {
	if token != "." {
		return nil, fmt.Errorf("invalid token")
	}

	op := NewMiscOperation(OP_DUMP, nil, TOKEN_DUMP)
	cnum := strings.Index(line, token)

	op.TokenStart().SetPostition(lnum+1, cnum+1)

	return op, nil
}

func TokenDo(token, line string, lnum int, blocks *BlockStack) (Operation, error) {
	if token != "do" {
		return nil, fmt.Errorf("invalid token")
	}

	block := NewMiscBlockOperation(OP_BLOCK, TOKEN_DO, TOKEN_END)
	blocks.Push(block)

	cnum := strings.Index(line, token)

	block.TokenStart().SetPostition(lnum+1, cnum+1)

	return nil, nil
}

func TokenIf(token, line string, lnum int, blocks *BlockStack) (Operation, error) {
	if token != "if" {
		return nil, fmt.Errorf("invalid token")
	}

	blockOperation := NewMiscBlockOperation(OP_BLOCK, TOKEN_IF, TOKEN_END)
	blocks.Push(blockOperation)

	cnum := strings.Index(line, token)

	blocks.Tail().TokenStart().SetPostition(lnum+1, cnum+1)

	return nil, nil
}

func TokenElse(token, line string, lnum int, blocks *BlockStack) (Operation, error) {
	if token != "else" {
		return nil, fmt.Errorf("invalid token")
	}

	cnum := strings.Index(line, token)

	elseBlock := NewMiscBlockOperation(OP_BLOCK, TOKEN_ELSE, TOKEN_END)
	elseBlock.TokenStart().SetPostition(lnum+1, cnum+1)

	block := blocks.Tail()

	if b := block.Tail(); block != nil {
		b.TokenEnd().SetPostition(lnum+1, cnum+1)
	}

	block.SetRefBlock(elseBlock)

	return nil, nil
}

func TokenEnd(token, line string, lnum int, blocks *BlockStack) (Operation, error) {
	if token != "end" {
		return nil, fmt.Errorf("invalid token")
	}

	block, err := blocks.Pop()
	if err != nil {
		return nil, err
	}

	cnum := strings.Index(line, token)

	if b := block.Tail(); b != nil {
		b.TokenEnd().SetPostition(lnum+1, cnum+1)
	}

	return block, nil
}
