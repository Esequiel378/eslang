package core

import (
	"fmt"
	"log"
	"reflect"
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
	Value(*Stack) (interface{}, error)
}

type MiscOperation struct {
	_type int
	value interface{}
}

func NewMiscOperation(operation int, value interface{}) Operation {
	if !REGISTERED_OPERATIONS[operation] {
		log.Fatal("invalid operation: ", operation)
	}

	return MiscOperation{
		_type: operation,
		value: value,
	}
}

func (mo MiscOperation) Type() int {
	return mo._type
}

func (mo MiscOperation) Value(stack *Stack) (interface{}, error) {
	return mo.value, nil
}

type BlockOperation interface {
	Type() int
	Value(*Stack) (interface{}, error)

	PushIntoBlocks(Operation)
	HasElseBlock() bool
	EnableElseBlock()
}

type MiscBlockOperation struct {
	_type        int
	block        *Program
	elseBlock    *Program
	hasElseBlock bool
}

func NewMiscBlockOperation() MiscBlockOperation {
	return MiscBlockOperation{
		_type:        OP_BLOCK,
		block:        &Program{},
		elseBlock:    &Program{},
		hasElseBlock: false,
	}
}

func (b *MiscBlockOperation) Type() int {
	return b._type
}

func (b *MiscBlockOperation) Value(stack *Stack) (interface{}, error) {
	return b.block, nil
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

type IfBlockOperation struct {
	_type        int
	block        *Program
	elseBlock    *Program
	hasElseBlock bool
}

func NewIfBlockOperation() IfBlockOperation {
	return IfBlockOperation{
		_type:        OP_BLOCK,
		block:        &Program{},
		elseBlock:    &Program{},
		hasElseBlock: false,
	}
}

func (b *IfBlockOperation) Type() int {
	return b._type
}

func (b *IfBlockOperation) Value(stack *Stack) (interface{}, error) {
	value, err := stack.Pop()

	if err != nil {
		return nil, err
	}

	truthy, ok := value.(int64)

	if !ok {
		return nil, fmt.Errorf(
			"error testing the truthy of %s with type %s",
			value,
			reflect.TypeOf(value),
		)
	}

	if truthy != 0 {
		return b.block, nil
	}

	if isNil(b.elseBlock) {
		return nil, nil
	}

	return b.elseBlock, nil
}

func (b *IfBlockOperation) PushIntoBlocks(operation Operation) {
	if b.HasElseBlock() {
		b.elseBlock.Push(operation)
	} else {
		b.block.Push(operation)
	}
}

func (b *IfBlockOperation) HasElseBlock() bool {
	return b.hasElseBlock
}

func (b *IfBlockOperation) EnableElseBlock() {
	b.hasElseBlock = true
}
