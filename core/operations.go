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
}

type MiscOperation struct {
	_type int
	value interface{}
}

func (mo MiscOperation) Type() int {
	return mo._type
}

func (mo MiscOperation) Value() interface{} {
	return mo.value
}

type BlockOperation struct {
	_type     int
	block     *Program
	elseBlock *Program
}

func (b BlockOperation) Type() int {
	return b._type
}

func (b BlockOperation) Value() interface{} {
	return b.block
}

func NewOP(operation int, value interface{}) Operation {
	if !REGISTERED_OPERATIONS[operation] {
		log.Fatal("invalid operation: ", operation)
	}

	return MiscOperation{
		_type: operation,
		value: value,
	}
}
