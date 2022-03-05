package core

import (
	"log"
)

const (
	OP_PUSH  = iota
	OP_PLUS  = iota
	OP_MINUS = iota
	OP_EQUAL = iota
	OP_BLOCK = iota
	OP_IF    = iota
	OP_DUMP  = iota
)

var BLOCK_ATTACH_OPERATIONS = map[int]bool{
	OP_IF: true,
}

var REGISTERED_OPERATIONS = map[int]bool{
	OP_PUSH:  true,
	OP_PLUS:  true,
	OP_MINUS: true,
	OP_EQUAL: true,
	OP_BLOCK: true,
	OP_DUMP:  true,
}

type Operation interface {
	Type() int
	Value() interface{}
	TokenStart() *Token
	TokenEnd() *Token
}

type MiscOperation struct {
	_type      int
	value      interface{}
	tokenStart *Token
	tokenEnd   *Token
}

func NewMiscOperation(operation int, value interface{}, token int) Operation {
	if !REGISTERED_OPERATIONS[operation] {
		log.Fatal("invalid operation: ", operation)
	}

	return &MiscOperation{
		_type: operation,
		value: value,
		tokenStart: &Token{
			line:  nil,
			col:   nil,
			token: token,
		},
		tokenEnd: &Token{
			line:  nil,
			col:   nil,
			token: token,
		},
	}
}

func (mo MiscOperation) Type() int {
	return mo._type
}

func (mo MiscOperation) Value() interface{} {
	return mo.value
}

func (mp MiscOperation) TokenStart() *Token {
	return mp.tokenStart
}

func (mp MiscOperation) TokenEnd() *Token {
	return mp.tokenEnd
}

type BlockOperation interface {
	// TODO: This are comming from Operation, is there a better way to do it?
	Type() int
	Value() interface{}
	TokenStart() *Token
	TokenEnd() *Token

	PushIntoBlocks(Operation)
	HasElseBlock() bool
	EnableElseBlock()
	ElseBlock() *Program
}

type MiscBlockOperation struct {
	_type        int
	block        *Program
	elseBlock    *Program
	hasElseBlock bool
	tokenStart   *Token
	tokenEnd     *Token
}

func NewMiscBlockOperation(operation int, tokenStart, tokenEnd int) BlockOperation {
	return &MiscBlockOperation{
		_type:        operation,
		block:        &Program{},
		elseBlock:    &Program{},
		hasElseBlock: false,
		tokenStart: &Token{
			line:  nil,
			col:   nil,
			token: tokenStart,
		},
		tokenEnd: &Token{
			line:  nil,
			col:   nil,
			token: tokenEnd,
		},
	}
}

func (b *MiscBlockOperation) Type() int {
	return b._type
}

func (b *MiscBlockOperation) Value() interface{} {
	return b.block
}

func (b *MiscBlockOperation) TokenStart() *Token {
	return b.tokenStart
}

func (b *MiscBlockOperation) TokenEnd() *Token {
	return b.tokenEnd
}

func (b *MiscBlockOperation) PushIntoBlocks(operation Operation) {
	if b.HasElseBlock() {
		b.elseBlock.Push(operation)
	} else {
		b.block.Push(operation)
	}
}

func (b *MiscBlockOperation) HasElseBlock() bool {
	return b.hasElseBlock
}

func (b *MiscBlockOperation) EnableElseBlock() {
	b.hasElseBlock = true
}

func (b *MiscBlockOperation) ElseBlock() *Program {
	return b.elseBlock
}
