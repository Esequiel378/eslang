package operations

// OPBlockIfElse struct    represents a  operation that starts a new block
type OPBlockIfElse struct {
	position   Position
	operations []Operation
	next       OperationLinkedBlocks
	isClosed   bool
}

// NewOPBlockIfElse function    creates a new OperationBlock
func NewOPBlockIfElse(position Position) *OPBlockIfElse {
	return &OPBlockIfElse{
		operations: []Operation{},
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
	return OP_BLOCK_IF_ELSE
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

	if lastOP.Type() == OP_PROGRAM {
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

// SetNext method    set the next block to start pushing operations
func (op *OPBlockIfElse) SetNext(position Position) {
	op.next = NewOPBlockIfElse(position)
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
