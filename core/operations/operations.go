package operations

// OPType int    returns the type of the operation
type OPType int

const (
	OP_PROGRAM OPType = iota
	OP_BLOCK_IF_ELSE
	OP_DUMP
	OP_PUSH_FLOAT
	OP_PUSH_INT
	OP_PUSH_STRING

	OP_TYPE_COUNT
)

// OP_TYPE_ALIASES map   is a map of OPType to their respective string representations
var OP_TYPE_ALIASES = map[OPType]string{
	OP_PROGRAM:       "OP_PROGRAM",
	OP_BLOCK_IF_ELSE: "OP_IF_ELSE",
	OP_DUMP:          "OP_DUMP",
	OP_PUSH_FLOAT:    "OP_PUSH_FLOAT",
	OP_PUSH_INT:      "OP_PUSH_INT",
	OP_PUSH_STRING:   "OP_PUSH_STRING",
}

// String method    returns the string representation of the operation
func (opType OPType) String() string {
	return OP_TYPE_ALIASES[opType]
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

// OperationBlock interface    represents a block of operations
type OperationBlock interface {
	// Position method    returns the position of the operation
	Position() Position
	// Type method    returns the type of the operation
	Type() OPType
	// Operations method    returns the operations of the block
	Operations() []Operation
	// IsEmpty method    returns true if the block is empty
	IsEmpty() bool
	// Push method    adds an operation to the block
	Push(operation Operation)
	// Last method    returns the last operation of the block
	LastOP() Operation
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
	// Close method    closes the last inner block
	CloseLastBlock()
	// IsClosed method    returns true if the block is closed
	IsClosed() bool
}
