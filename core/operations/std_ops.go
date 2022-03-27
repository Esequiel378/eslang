package operations

// OPDump struct  Óòß  represents a  operation that dumps the stack
type OPDump struct {
	position Position
}

// NewOPDump function  Óòß  creates a new OperationDump
func NewOPDump(position Position) Operation {
	return OPDump{
		position: position,
	}
}

// Position method  Óòß  returns the position of the operation
func (op OPDump) Position() Position {
	return op.position
}

// Type method  Óòß  returns the type of the operation
func (op OPDump) Type() OPType {
	return OP_DUMP
}
