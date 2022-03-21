package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type (
	TokenType int
	// TokenHandler type    function definition to implement new token handlers
	TokenHandler func(token, line string, lnum int, program *Program) (bool, error)
)

const (
	TOKEN_DO TokenType = iota
	TOKEN_DUMP
	TOKEN_ELSE
	TOKEN_END
	TOKEN_EQUAL
	TOKEN_IF
	TOKEN_INT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_PUSH_FLOAT
	TOKEN_PUSH_INT
	TOKEN_PUSH_STR
	TOKEN_VAR
	TOKEN_VAR_READ
	TOKEN_VAR_WRITE
)

var REGISTERED_TOKENS = map[TokenType]TokenHandler{
	TOKEN_DO:         TokenDo,
	TOKEN_DUMP:       TokenDump,
	TOKEN_ELSE:       TokenElse,
	TOKEN_END:        TokenEnd,
	TOKEN_EQUAL:      TokenEqual,
	TOKEN_IF:         TokenIf,
	TOKEN_INT:        TokenInt,
	TOKEN_MINUS:      TokenMinus,
	TOKEN_PLUS:       TokenPlus,
	TOKEN_PUSH_FLOAT: TokenPushFloat,
	TOKEN_PUSH_INT:   TokenPushInt,
	TOKEN_PUSH_STR:   TokenPushStr,
	TOKEN_VAR:        TokenVar,
	TOKEN_VAR_READ:   TokenVarRead,
	TOKEN_VAR_WRITE:  TokenVarWrite,
}

var TOKEN_ALIASES = map[TokenType]string{
	TOKEN_DO:         "DO",
	TOKEN_DUMP:       "DUMP",
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
}

var (
	IS_INT            = regexp.MustCompile(`^\d+$`)
	IS_FLOAT          = regexp.MustCompile(`^\d+\.\d+$`)
	IS_STR            = regexp.MustCompile(`^".+"$`)
	IS_VALID_VARIABLE = regexp.MustCompile(`^[a-zA-Z]+$`)
)

type Token struct {
	line  int
	col   int
	token TokenType
}

func (t *Token) Token() TokenType {
	return t.token
}

func (t *Token) TokenAlias() string {
	return TOKEN_ALIASES[t.token]
}

func (t *Token) Position() (int, int) {
	return t.line, t.col
}

func (t *Token) SetPostition(line, col int) {
	t.line = line
	t.col = col
}

func TokenPushInt(token, line string, lnum int, program *Program) (bool, error) {
	if !IS_INT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	opValue := NewOperationValue().SetInt(value)
	op := NewOperation(OP_PUSH_INT, opValue, TOKEN_PUSH_INT, TOKEN_PUSH_INT)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenPushFloat(token, line string, lnum int, program *Program) (bool, error) {
	if !IS_FLOAT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	opValue := NewOperationValue().SetFloat(value)
	op := NewOperation(OP_PUSH_FLOAT, opValue, TOKEN_PUSH_FLOAT, TOKEN_PUSH_FLOAT)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenPushStr(token, line string, lnum int, program *Program) (bool, error) {
	if !IS_STR.MatchString(token) {
		return false, nil
	}

	opValue := NewOperationValue().SetStr(token)
	op := NewOperation(OP_PUSH_STR, opValue, TOKEN_PUSH_STR, TOKEN_PUSH_STR)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenInt(token, line string, lnum int, program *Program) (bool, error) {
	if token != "int" {
		return false, nil
	}

	lastOP := program.Last()

	if lastOP == nil || lastOP.Type() != OP_VAR {
		cnum := strings.Index(line, token)
		return true, fmt.Errorf("error: using type `int` out of context at line %d:%d", lnum+1, cnum+1)
	}

	lastOP.Value().SetType(Int)

	return true, nil
}

func TokenPlus(token, line string, lnum int, program *Program) (bool, error) {
	if token != "+" {
		return false, nil
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_MOP, opValue, TOKEN_PLUS, TOKEN_PLUS)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenVar(token, line string, lnum int, program *Program) (bool, error) {
	if !IS_VALID_VARIABLE.MatchString(token) || RESERVED_WORDS[token] {
		return false, nil
	}

	opValue := NewOperationValue().SetName(token)
	op := NewOperation(OP_VAR, opValue, TOKEN_VAR, TOKEN_VAR)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	varOp, found := program.GetVariable(token)

	if found {
		t := varOp.Value().Type()
		op.Value().SetType(t)
	} else {
		program.SetVariable(token, op)
	}

	program.Push(op)

	return true, nil
}

func TokenVarRead(token, line string, lnum int, program *Program) (bool, error) {
	if token != "," {
		return false, nil
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_VAR_READ, opValue, TOKEN_VAR_READ, TOKEN_VAR_READ)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenVarWrite(token, line string, lnum int, program *Program) (bool, error) {
	if token != "." {
		return false, nil
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_VAR_WRITE, opValue, TOKEN_VAR_WRITE, TOKEN_VAR_WRITE)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenMinus(token, line string, lnum int, program *Program) (bool, error) {
	if token != "-" {
		return false, nil
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_MOP, opValue, TOKEN_MINUS, TOKEN_MINUS)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenEqual(token, line string, lnum int, program *Program) (bool, error) {
	if token != "=" {
		return false, nil
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_MOP, opValue, TOKEN_EQUAL, TOKEN_EQUAL)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenDump(token, line string, lnum int, program *Program) (bool, error) {
	if token != "dump" {
		return false, nil
	}

	opValue := NewOperationValue()
	op := NewOperation(OP_DUMP, opValue, TOKEN_DUMP, TOKEN_DUMP)

	cnum := strings.Index(line, token)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenDo(token, line string, lnum int, program *Program) (bool, error) {
	if token != "do" {
		return false, nil
	}

	cnum := strings.Index(line, token)

	opValue := NewOperationValue()
	opValue.Block().TokenStart().SetPostition(lnum+1, cnum+1)

	op := NewOperation(OP_BLOCK, opValue, TOKEN_DO, TOKEN_END)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenIf(token, line string, lnum int, program *Program) (bool, error) {
	if token != "if" {
		return false, nil
	}

	cnum := strings.Index(line, token)

	opValue := NewOperationValue()
	opValue.Block().TokenStart().SetPostition(lnum+1, cnum+1)

	op := NewOperation(OP_BLOCK, opValue, TOKEN_IF, TOKEN_END)
	op.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Push(op)

	return true, nil
}

func TokenElse(token, line string, lnum int, program *Program) (bool, error) {
	if token != "else" {
		return false, nil
	}

	cnum := strings.Index(line, token)

	if err := program.CloseLastBlock(lnum+1, cnum+1); err != nil {
		return true, err
	}

	block := NewBlock(TOKEN_ELSE, TOKEN_END)
	block.TokenStart().SetPostition(lnum+1, cnum+1)

	program.Last().Value().Block().SetNext(block)

	return true, nil
}

func TokenEnd(token, line string, lnum int, program *Program) (bool, error) {
	if token != "end" {
		return false, nil
	}

	cnum := strings.Index(line, token)

	if err := program.CloseLastBlock(lnum+1, cnum+1); err != nil {
		return true, err
	}

	return true, nil
}
