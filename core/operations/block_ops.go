package operations

// OPBlockIfElse struct  Óòß  represents a  operation that starts a new block
type OPBlockIfElse struct {
	position   Position
	operations []Operation
	next       OperationLinkedBlocks
	isClosed   bool
}

// NewOPBlockIfElse function  Óòß  creates a new OperationBlock
func NewOPBlockIfElse(operations []Operation, position Position) *OPBlockIfElse {
	return &OPBlockIfElse{
		operations: operations,
		position:   position,
		next:       nil,
		isClosed:   false,
	}
}

// Position method  Óòß  returns the position of the operation
func (op *OPBlockIfElse) Position() Position {
	return op.position
}

// Type method  Óòß  returns the type of the operation
func (op *OPBlockIfElse) Type() OPType {
	return OP_BLOCK_IF_ELSE
}

// Operations method  Óòß  returns the operations of the block
func (op *OPBlockIfElse) Operations() []Operation {
	return op.operations
}

// IsEmpty method  Óòß  returns true if the block is empty
func (op *OPBlockIfElse) IsEmpty() bool {
	return len(op.operations) == 0
}

// Push method  Óòß  adds an operation to the block
func (op *OPBlockIfElse) Push(operation Operation) {
	op.operations = append(op.operations, operation)
}

// LastOP method  Óòß  returns the last operation of the block
func (op *OPBlockIfElse) LastOP() Operation {
	return op.operations[len(op.operations)-1]
}

// CloseLastBlock method  Óòß  closes the last inner block
func (op *OPBlockIfElse) CloseLastBlock() {
	if op.HasNext() {
		op.Next().CloseLastBlock()
	}

	lastOP := op.LastOP()

	if lastOP.Type() == OP_PROGRAM {
		block := lastOP.(OperationLinkedBlocks)
		block.CloseLastBlock()
		return
	}

	op.isClosed = true
}

// IsClosed method  Óòß  returns true if the block is closed
func (op *OPBlockIfElse) IsClosed() bool {
	return op.isClosed
}

// HasNext  Óòß  returns true if the block has a linked block
func (op *OPBlockIfElse) HasNext() bool {
	return op.next != nil
}

// Next method  Óòß  returns the next block
func (op *OPBlockIfElse) Next() OperationLinkedBlocks {
	return op.next
}

// Last method  Óòß  returns the last linked block
func (op *OPBlockIfElse) LastBlock() OperationLinkedBlocks {
	if !op.HasNext() {
		return op.Next()
	}

	return op.Next()
}
