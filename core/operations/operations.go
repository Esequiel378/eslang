package operations

// OPType int    returns the type of the operation
type OPType int

const (
	OP_PROGRAM OPType = iota
	// blocks
	OP_BLOCK_IF_ELSE
	OP_BLOCK_WHILE

	// std
	OP_DUMP

	// push
	OP_PUSH_FLOAT
	OP_PUSH_INT
	OP_PUSH_STRING
	OP_PUSH_BOOL

	// variables
	OP_VARIABLE
	OP_VARIABLE_WRITE

	// arithmetic operators
	OP_OPERATOR_ADD
	OP_OPERATOR_SUB
	OP_OPERATOR_MUL
	OP_OPERATOR_DIV
	OP_OPERATOR_MOD

	// realtional operators
	OP_R_OPERATOR_EQUAL
	OP_R_OPERATOR_NOT_EQUAL
	OP_R_OPERATOR_LESS_THAN
	OP_R_OPERATOR_LESS_THAN_OR_EQUAL
	OP_R_OPERATOR_GREATER_THAN
	OP_R_OPERATOR_GREATER_THAN_OR_EQUAL

	OP_TYPE_COUNT
)

// OP_TYPE_ALIASES map   is a map of OPType to their respective string representations
var OP_TYPE_ALIASES = map[OPType]string{
	OP_PROGRAM: "OP_PROGRAM",

	// blocks
	OP_BLOCK_IF_ELSE: "OP_BLOCK_IF_ELSE",
	OP_BLOCK_WHILE:   "OP_BLOCK_WHILE",

	// std
	OP_DUMP: "OP_DUMP",

	// push
	OP_PUSH_FLOAT:  "OP_PUSH_FLOAT",
	OP_PUSH_INT:    "OP_PUSH_INT",
	OP_PUSH_STRING: "OP_PUSH_STRING",
	OP_PUSH_BOOL:   "OP_PUSH_BOOL",

	// variables
	OP_VARIABLE:       "OP_VARIABLE",
	OP_VARIABLE_WRITE: "OP_VARIABLE_WRITE",

	// arithmetic operators
	OP_OPERATOR_ADD: "OP_OPERATOR_ADD",
	OP_OPERATOR_SUB: "OP_OPERATOR_SUB",
	OP_OPERATOR_MUL: "OP_OPERATOR_MUL",
	OP_OPERATOR_DIV: "OP_OPERATOR_DIV",
	OP_OPERATOR_MOD: "OP_OPERATOR_MOD",

	// relational operators
	OP_R_OPERATOR_EQUAL:                 "OP_R_OPERATOR_EQUAL",
	OP_R_OPERATOR_NOT_EQUAL:             "OP_R_OPERATOR_NOT_EQUAL",
	OP_R_OPERATOR_LESS_THAN:             "OP_R_OPERATOR_LESS_THAN",
	OP_R_OPERATOR_LESS_THAN_OR_EQUAL:    "OP_R_OPERATOR_LESS_THAN_EQUAL",
	OP_R_OPERATOR_GREATER_THAN:          "OP_R_OPERATOR_GREATER_THAN",
	OP_R_OPERATOR_GREATER_THAN_OR_EQUAL: "OP_R_OPERATOR_GREATER_THAN_EQUAL",
}

// String method    returns the string representation of the operation
func (opType OPType) String() string {
	return OP_TYPE_ALIASES[opType]
}

// IsBlock method    returns true if the operation is a block
func (opType OPType) IsBlock() bool {
	switch opType {
	case OP_BLOCK_IF_ELSE, OP_BLOCK_WHILE:
		return true
	}

	return false
}

// Position struct    represents the operation position in the source code
type Position struct {
	line   int
	column int
	file   string
}

// NewPosition function    creates a new Position
func NewPosition(line, column int, file string) Position {
	return Position{
		line:   line,
		column: column,
		file:   file,
	}
}

// File method    returns the file where the operation is located
func (p Position) File() string {
	return p.file
}

// Ruler method    returns the line and column of the operation
func (p Position) Ruler() (int, int) {
	return p.line, p.column
}

// Operation interface    represents a single operation in the program
type Operation interface {
	// Position method    returns the position of the operation
	Position() Position
	// Type method    returns the type of the operation
	Type() OPType
}

// MiscOperation struct    represents a operation that only holds a type and a position
type MiscOperation struct {
	position Position
	opType   OPType
}

// NewMiscOperation function    creates a new MiscOperation
func NewMiscOperation(position Position, opType OPType) Operation {
	return MiscOperation{
		position: position,
		opType:   opType,
	}
}

// Position method    returns the position of the operation
func (op MiscOperation) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op MiscOperation) Type() OPType {
	return op.opType
}

// OperationBlock interface    represents a block of operations
type OperationBlock interface {
	// Position method    returns the position of the operation
	Position() Position
	// Type method    returns the type of the operation
	Type() OPType
	// Program method    returns the program of the block
	Program() *Program
	// IsEmpty method    returns true if the block is empty
	IsEmpty() bool
	// Push method    adds an operation to the block
	Push(operation Operation) error
	// Last method    returns the last operation of the block
	LastOP() Operation
	// IsClosed method    returns true if the block is closed
	IsClosed() bool
	// CloseBlock method    closes the last inner block
	CloseBlock()
}

// OperationLinkedBlocks interface    represents a block of operations that can be linked to other blocks
type OperationLinkedBlocks interface {
	// HasNext method    returns true if the block has a next block attached
	HasNext() bool
	// SetNext method    set the next block to start pushing operations
	SetNext(Position)
	// Next method    returns the next block attached to the current block
	Next() OperationLinkedBlocks
	// LastBlock method    returns the last block attached to the current block
	LastBlock() OperationLinkedBlocks
	// CloseLastBlock method    closes the last inner block
	CloseLastBlock()
}
