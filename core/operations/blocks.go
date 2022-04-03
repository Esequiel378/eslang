// TODO: make variables scoped to blocks
package operations

import "fmt"

// OPBlockIfElse struct    represents a  operation that starts a new block
type OPBlockIfElse struct {
	position Position
	program  *Program
	next     OperationLinkedBlocks
	isOpen   bool
}

// NewOPBlockIfElse function    creates a new OperationBlock
func NewOPBlockIfElse(program *Program, position Position) OperationLinkedBlocks {
	return &OPBlockIfElse{
		program:  program,
		position: position,
		next:     nil,
		isOpen:   true,
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
	if op.isOpen {
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

// LastNestedBlock method    returns the last nested block
func (op *OPBlockIfElse) LastNestedBlock() OperationBlock {
	block := op.Program().LastOP()

	if block == nil || !block.Type().IsBlock() {
		return op
	}

	b := block.(OperationBlock)

	if b.IsOpen() {
		return b.LastNestedBlock()
	}

	return op
}

// CloseLastBlock method    closes the last inner block
func (op *OPBlockIfElse) CloseLastBlock() {
	if op.HasNext() {
		op.Next().CloseLastBlock()
	}

	lastOP := op.LastOP()

	if lastOP != nil && lastOP.Type() == OP_PROGRAM {
		block := lastOP.(OperationLinkedBlocks)

		if block.LastBlock().IsOpen() {
			block.CloseLastBlock()
			return

		}
	}

	op.isOpen = false
}

// CloseBlock method    closes the block
func (op *OPBlockIfElse) CloseBlock() {
	op.CloseLastBlock()
}

// IsOpen method    returns true if the block is closed
func (op *OPBlockIfElse) IsOpen() bool {
	open := op.isOpen

	if op.HasNext() {
		open = op.Next().IsOpen()
	}

	return open
}

// HasNext    returns true if the block has a linked block
func (op *OPBlockIfElse) HasNext() bool {
	return op.next != nil
}

// SetNext method    set the next block to start pushing operations
func (op *OPBlockIfElse) SetNext(block OperationLinkedBlocks) {
	op.next = block
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
	position        Position
	program         *Program
	isOpen          bool
	condition       []Operation
	isConditionOpen bool
}

// NewOPBlockWhile function    creates a new OPBlockWhile
func NewOPBlockWhile(program *Program, position Position) OperationLoop {
	return &OPBlockWhile{
		program:         program,
		position:        position,
		condition:       make([]Operation, 0),
		isOpen:          true,
		isConditionOpen: true,
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
	if op.isConditionOpen {
		op.condition = append(op.condition, operation)
		return nil
	}

	if op.isOpen {
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

// LastNestedBlock method    returns the last nested block
// TODO: somehow check when the block condition is not closed and rise error
func (op *OPBlockWhile) LastNestedBlock() OperationBlock {
	block := op.Program().LastOP()

	if block == nil || !block.Type().IsBlock() {
		return op
	}

	b := block.(OperationBlock)

	if b.IsOpen() {
		return b.LastNestedBlock()
	}

	return op
}

// IsOpen method    returns true if the block is closed
func (op *OPBlockWhile) IsOpen() bool {
	return op.isOpen
}

// CloseBlock method    closes the block
func (op *OPBlockWhile) CloseBlock() {
	op.isOpen = false
}

// ConditionBlock method    returns the condition block
func (op *OPBlockWhile) ConditionBlock() []Operation {
	return op.condition
}

// IsConditionBlockOpen method    returns true if the condition block is open
func (op *OPBlockWhile) IsConditionBlockOpen() bool {
	return op.isConditionOpen
}

// CloseConditionBlock method    closes the condition block
func (op *OPBlockWhile) CloseConditionBlock() {
	op.isConditionOpen = false
}
