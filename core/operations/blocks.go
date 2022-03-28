package operations

import "fmt"

// OPBlockIfElse struct    represents a  operation that starts a new block
type OPBlockIfElse struct {
	position Position
	program  *Program
	next     OperationLinkedBlocks
	isClosed bool
}

// NewOPBlockIfElse function    creates a new OperationBlock
func NewOPBlockIfElse(position Position) *OPBlockIfElse {
	return &OPBlockIfElse{
		program:  &Program{},
		position: position,
		next:     nil,
		isClosed: false,
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

// Program method    returns the program of the block
func (op *OPBlockIfElse) Program() *Program {
	return op.program
}

// IsEmpty method    returns true if the block is empty
func (op *OPBlockIfElse) IsEmpty() bool {
	return op.program.IsEmpty()
}

// Push method    adds an operation to the block
func (op *OPBlockIfElse) Push(operation Operation) error {
	if !op.isClosed {
		err := op.Program().Push(operation)
		return err
	}

	if op.HasNext() {
		next := op.Next().(*OPBlockIfElse)
		err := next.Push(operation)

		return err
	}

	return fmt.Errorf("block is closed")
}

// LastOP method    returns the last operation of the block
func (op *OPBlockIfElse) LastOP() Operation {
	if op.IsEmpty() {
		return nil
	}

	return op.program.LastOP()
}

// CloseLastBlock method    closes the last inner block
func (op *OPBlockIfElse) CloseLastBlock() {
	if op.HasNext() {
		op.Next().CloseLastBlock()
	}

	lastOP := op.LastOP()

	if lastOP != nil && lastOP.Type() == OP_PROGRAM {
		block := lastOP.(OperationLinkedBlocks)
		block.CloseLastBlock()
		return
	}

	op.isClosed = true
}

// CloseBlock method    closes the block
func (op *OPBlockIfElse) CloseBlock() {
	op.CloseLastBlock()
}

// IsClosed method    returns true if the block is closed
func (op *OPBlockIfElse) IsClosed() bool {
	closed := op.isClosed

	if op.HasNext() {
		next := op.Next().(*OPBlockIfElse)
		closed = next.IsClosed()
	}

	return closed
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

// OPBlockWhile struct    represents a  operation that starts a while block
type OPBlockWhile struct {
	position Position
	program  *Program
	isClosed bool
}

// NewOPBlockWhile function    creates a new OPBlockWhile
func NewOPBlockWhile(position Position) *OPBlockWhile {
	return &OPBlockWhile{
		program:  &Program{},
		position: position,
		isClosed: false,
	}
}

// Position method    returns the position of the operation
func (op *OPBlockWhile) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op *OPBlockWhile) Type() OPType {
	return OP_BLOCK_WHILE
}

// Program method    returns the program of the block
func (op *OPBlockWhile) Program() *Program {
	return op.program
}

// IsEmpty method    returns true if the block is empty
func (op *OPBlockWhile) IsEmpty() bool {
	return op.program.IsEmpty()
}

// Push method    adds an operation to the block
func (op *OPBlockWhile) Push(operation Operation) error {
	if !op.isClosed {
		err := op.Program().Push(operation)
		return err
	}

	return fmt.Errorf("block is closed")
}

// LastOP method    returns the last operation of the block
func (op *OPBlockWhile) LastOP() Operation {
	if op.IsEmpty() {
		return nil
	}

	return op.program.LastOP()
}

// IsClosed method    returns true if the block is closed
func (op *OPBlockWhile) IsClosed() bool {
	return op.isClosed
}

// CloseBlock method    closes the block
func (op *OPBlockWhile) CloseBlock() {
	op.isClosed = true
}
