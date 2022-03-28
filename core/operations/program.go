package operations

import "fmt"

// TODO: check if there is any open block at the end of the program
// TODO: and end token to close blocks

// Program struct    represents a program block (entry point of the code)
type Program struct {
	position   Position
	operations []Operation
	variables  map[string]*OPVariable
}

// NewProgram function    creates a new program block
func NewProgram(filename string) *Program {
	position := NewPosition(0, 0, filename)

	return &Program{
		position:   position,
		operations: make([]Operation, 0),
		variables:  make(map[string]*OPVariable),
	}
}

// SetVariable method    sets a variable in the program
func (p *Program) SetVariable(name string, value *OPVariable) {
	p.variables[name] = value
}

// GetVariable method    returns the variable of the program
func (p *Program) GetVariable(name string) (*OPVariable, bool) {
	variable, found := p.variables[name]

	return variable, found
}

// PeekTwo method    returns the two last operations of the program without removing them
func (p *Program) PeekTwo() (lhs Operation, rhs Operation, err error) {
	if len(p.operations) < 2 {
		err = fmt.Errorf("not enough operations in program")
		return
	}

	lhs = p.operations[len(p.operations)-2]
	rhs = p.operations[len(p.operations)-1]

	return
}

// Position method    returns the position of the operation
func (op *Program) Position() Position {
	return op.position
}

// Type method    returns the type of the operation
func (op *Program) Type() OPType {
	return OP_PROGRAM
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
func (op *Program) Push(operation Operation) error {
	lastOP := op.LastOP()

	// Check if the last operation is a block
	if lastOP != nil && lastOP.Type().IsBlock() {
		block := lastOP.(OperationBlock)

		// Push the operation to the block is it's not close
		if !block.IsClosed() {
			err := block.Push(operation)
			return err
		}
	}

	// Check if the last operation is a variable declaration
	if lastOP != nil && lastOP.Type() == OP_VARIABLE {
		variable := lastOP.(*OPVariable)
		name := variable.Name()

		// Check if the variable is not initialized and return an error
		if _, found := op.GetVariable(name); !found {
			line, column := variable.Position().Ruler()
			file := variable.Position().File()

			return fmt.Errorf("uninitialised variable %s in %s:%d:%d", name, file, line, column)
		}
	}

	// Push the operation to the program
	op.operations = append(op.operations, operation)

	return nil
}

// LastOP method    returns the last operation of the program
func (op *Program) LastOP() Operation {
	if op.IsEmpty() {
		return nil
	}

	return op.operations[len(op.operations)-1]
}
