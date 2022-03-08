package core

import (
	"log"
)

// TODO: Use custom int type to make it more type safe

// type OPTYPE in
// const OP_BLOCK OPTYPE = iota

// Same for tokens and any iotas

const (
	OP_BLOCK = iota
	OP_DUMP  = iota
	OP_ELSE  = iota
	OP_EQUAL = iota
	OP_IF    = iota
	OP_MINUS = iota
	OP_PLUS  = iota
	OP_PUSH  = iota
)

var REGISTERED_OPERATIONS = map[int]bool{
	OP_BLOCK: true,
	OP_DUMP:  true,
	OP_ELSE:  true,
	OP_EQUAL: true,
	OP_IF:    true,
	OP_MINUS: true,
	OP_PLUS:  true,
	OP_PUSH:  true,
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

	Block() *Program
	HasRefBlock() bool
	SetRefBlock(BlockOperation)
	RefBlock() BlockOperation
	Tail() BlockOperation
}

type MiscBlockOperation struct {
	_type      int
	block      *Program
	refBlock   BlockOperation
	tokenStart *Token
	tokenEnd   *Token
}

func NewMiscBlockOperation(operation int, tokenStart, tokenEnd int) BlockOperation {
	if !REGISTERED_OPERATIONS[operation] {
		log.Fatal("invalid operation: ", operation)
	}

	return &MiscBlockOperation{
		_type:    operation,
		block:    &Program{},
		refBlock: nil,
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

func (b *MiscBlockOperation) Block() *Program {
	return b.block
}

func (b *MiscBlockOperation) HasRefBlock() bool {
	return b.refBlock != nil
}

func (b *MiscBlockOperation) SetRefBlock(block BlockOperation) {
	b.refBlock = block
}

func (b *MiscBlockOperation) RefBlock() BlockOperation {
	return b.refBlock
}

func (b *MiscBlockOperation) Tail() BlockOperation {
	if b.HasRefBlock() {
		return b.RefBlock().Tail()
	}

	return b
}
