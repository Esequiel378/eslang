package core

// Program struct    represents a program block (entry point of the code)
type Program struct {
	position   Position
	operations []Operation
}

// NewProgram function    creates a new program block
func NewProgram(filename string) *Program {
	position := NewPosition(0, 0, filename)

	return &Program{
		position:   position,
		operations: make([]Operation, 0),
	}
}

// Position method    returns the position of the operation
func (op *Program) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op *Program) Type() OPType {
	return OP_BLOCK
}

// Operations method    returns the operations of the program
func (op *Program) Operations() []Operation {
	return op.operations
}

// IsEmpty method    returns true if the operation is empty
func (op *Program) IsEmpty() bool {
	return len(op.operations) == 0
}

// Push method    pushes an operation to the program
func (op *Program) Push(operation Operation) {
	op.operations = append(op.operations, operation)
}

// LastOP method    returns the last operation of the program
func (op *Program) LastOP() Operation {
	return op.operations[len(op.operations)-1]
}
