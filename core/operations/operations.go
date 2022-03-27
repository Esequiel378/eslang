package operations

type OPType int

const (
	OP_BLOCK OPType = iota
	OP_DUMP
	OP_PUSH_FLOAT
	OP_PUSH_INT
	OP_PUSH_STRING

	OP_TYPE_COUNT
)

var OP_TYPE_ALIASES = map[OPType]string{
	OP_BLOCK:       "OP_BLOCK",
	OP_DUMP:        "OP_DUMP",
	OP_PUSH_FLOAT:  "OP_PUSH_FLOAT",
	OP_PUSH_INT:    "OP_PUSH_INT",
	OP_PUSH_STRING: "OP_PUSH_STRING",
}

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

// OPPushInt struct    represents a  operation that pushes an integer onto the stack
type OPPushInt struct {
	position Position
	value    int64
}

// NewOPPushInt function    creates a new OperationPushInt
func NewOPPushInt(value int64, position Position) Operation {
	return OPPushInt{
		value:    value,
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OPPushInt) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OPPushInt) Type() OPType {
	return OP_PUSH_INT
}

// Value method    returns the value of the operation
func (op OPPushInt) Value() int64 {
	return op.value
}

// OPPushFloat struct    represents a  operation that pushes a float onto the stack
type OPPushFloat struct {
	position Position
	value    float64
}

// NewOPPushFloat function  
func NewOPPushFloat(value float64, position Position) Operation {
	return OPPushFloat{
		value:    value,
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OPPushFloat) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OPPushFloat) Type() OPType {
	return OP_PUSH_FLOAT
}

// Value method    returns the value of the operation
func (op OPPushFloat) Value() float64 {
	return op.value
}

// OPPushString struct    represents a  operation that pushes a string onto the stack
type OPPushString struct {
	position Position
	value    string
}

// NewOPPushString function    creates a new OperationPushStr
func NewOPPushString(value string, position Position) Operation {
	return OPPushString{
		value:    value,
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OPPushString) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OPPushString) Type() OPType {
	return OP_PUSH_STRING
}

// Value method    returns the value of the operation
func (op OPPushString) Value() string {
	return op.value
}

// OPDump struct    represents a  operation that dumps the stack
type OPDump struct {
	position Position
}

// NewOPDump function    creates a new OperationDump
func NewOPDump(position Position) Operation {
	return OPDump{
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OPDump) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OPDump) Type() OPType {
	return OP_DUMP
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
	// Next method    returns the next block attached to the current block
	Next() OperationLinkedBlocks
	// LastBlock method    returns the last block attached to the current block
	LastBlock() OperationLinkedBlocks
	// Close method    closes the last inner block
	CloseLastBlock()
	// IsClosed method    returns true if the block is closed
	IsClosed() bool
}

// OPBlockIfElse struct    represents a  operation that starts a new block
type OPBlockIfElse struct {
	position   Position
	operations []Operation
	next       OperationLinkedBlocks
	isClosed   bool
}

// NewOPBlockIfElse function    creates a new OperationBlock
func NewOPBlockIfElse(operations []Operation, position Position) *OPBlockIfElse {
	return &OPBlockIfElse{
		operations: operations,
		position:   position,
		next:       nil,
		isClosed:   false,
	}
}

// Position method    returns the position of the operation
func (op *OPBlockIfElse) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op *OPBlockIfElse) Type() OPType {
	return OP_BLOCK
}

// Operations method    returns the operations of the block
func (op *OPBlockIfElse) Operations() []Operation {
	return op.operations
}

// IsEmpty method    returns true if the block is empty
func (op *OPBlockIfElse) IsEmpty() bool {
	return len(op.operations) == 0
}

// Push method    adds an operation to the block
func (op *OPBlockIfElse) Push(operation Operation) {
	op.operations = append(op.operations, operation)
}

// LastOP method    returns the last operation of the block
func (op *OPBlockIfElse) LastOP() Operation {
	return op.operations[len(op.operations)-1]
}

// CloseLastBlock method    closes the last inner block
func (op *OPBlockIfElse) CloseLastBlock() {
	if op.HasNext() {
		op.Next().CloseLastBlock()
	}

	lastOP := op.LastOP()

	if lastOP.Type() == OP_BLOCK {
		block := lastOP.(OperationLinkedBlocks)
		block.CloseLastBlock()
		return
	}

	op.isClosed = true
}

// IsClosed method    returns true if the block is closed
func (op *OPBlockIfElse) IsClosed() bool {
	return op.isClosed
}

// HasNext    returns true if the block has a linked block
func (op *OPBlockIfElse) HasNext() bool {
	return op.next != nil
}

// Next method    returns the next block
func (op *OPBlockIfElse) Next() OperationLinkedBlocks {
	return op.next
}

// Last method    returns the last linked block
func (op *OPBlockIfElse) LastBlock() OperationLinkedBlocks {
	if !op.HasNext() {
		return op.Next()
	}

	return op.Next()
}
