package core

import (
	"log"
)

const (
	OP_PUSH  = iota
	OP_PLUS  = iota
	OP_MINUS = iota
	OP_EQUAL = iota
	OP_DUMP  = iota
)

type Operation struct {
	Type  int
	Value interface{}
}

var VALID_OPERATIONS = map[int]bool{
	OP_PUSH:  true,
	OP_PLUS:  true,
	OP_MINUS: true,
	OP_EQUAL: true,
	OP_DUMP:  true,
}

func NewOP(operation int, value interface{}) *Operation {
	if !VALID_OPERATIONS[operation] {
		log.Fatal("invalid operation: ", operation)
	}

	return &Operation{
		Type:  operation,
		Value: value,
	}
}
