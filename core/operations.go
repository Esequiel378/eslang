package core

type OPType int

const (
	OP_DUMP OPType = iota
	OP_PUSH_FLOAT
	OP_PUSH_INT
	OP_PUSH_STRING

	OP_TYPE_COUNT
)

var OP_TYPE_ALIASES = map[OPType]string{
	OP_DUMP:        "OP_DUMP",
	OP_PUSH_FLOAT:  "OP_PUSH_FLOAT",
	OP_PUSH_INT:    "OP_PUSH_INT",
	OP_PUSH_STRING: "OP_PUSH_STRING",
}

func (opType OPType) String() string {
	if alias, ok := OP_TYPE_ALIASES[opType]; ok {
		return alias
	}

	return "-unknown-"
}

// Position struct    represents the operation position in the source code
type Position struct {
	line   int
	column int
	file   string
}

// NewPosition function    creates a new Position
func NewPosition(line, column int, file string) *Position {
	return &Position{
		line:   line,
		column: column,
		file:   file,
	}
}

// File method    returns the file where the operation is located
func (p *Position) File() string {
	return p.file
}

// Ruler method    returns the line and column of the operation
func (p Position) Ruler() (int, int) {
	return p.line, p.column
}

// Operation interface    represents a single operation in the program
type Operation interface {
	Position() *Position
	Type() OPType
}

// OperationPushInt struct    represents a  operation that pushes an integer onto the stack
type OperationPushInt struct {
	position *Position
	value    int64
}

// NewOperationInt function    creates a new OperationPushInt
func NewOperationInt(value int64, position *Position) Operation {
	return OperationPushInt{
		value:    value,
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OperationPushInt) Position() *Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OperationPushInt) Type() OPType {
	return OP_PUSH_INT
}

// Value method    returns the value of the operation
func (op OperationPushInt) Value() int64 {
	return op.value
}

// OperationPushFloat struct    represents a  operation that pushes a float onto the stack
type OperationPushFloat struct {
	position *Position
	value    float64
}

// NewOperationFloat function  
func NewOperationFloat(value float64, position *Position) Operation {
	return OperationPushFloat{
		value:    value,
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OperationPushFloat) Position() *Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OperationPushFloat) Type() OPType {
	return OP_PUSH_FLOAT
}

// Value method    returns the value of the operation
func (op OperationPushFloat) Value() float64 {
	return op.value
}

// OperationPushString struct    represents a  operation that pushes a string onto the stack
type OperationPushString struct {
	position *Position
	value    string
}

// NewOperationString function    creates a new OperationPushStr
func NewOperationString(value string, position *Position) Operation {
	return OperationPushString{
		value:    value,
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OperationPushString) Position() *Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OperationPushString) Type() OPType {
	return OP_PUSH_STRING
}

// Value method    returns the value of the operation
func (op OperationPushString) Value() string {
	return op.value
}

// OperationDump struct    represents a  operation that dumps the stack
type OperationDump struct {
	position *Position
}

// NewOperationDump function    creates a new OperationDump
func NewOperationDump(position *Position) Operation {
	return OperationDump{
		position: position,
	}
}

// Position method    returns the position of the operation
func (op OperationDump) Position() *Position {
	return op.position
}

// Type method    returns the type of the operation
func (op OperationDump) Type() OPType {
	return OP_DUMP
}
