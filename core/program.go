package core

// Program struct    A program is a collection of operations
type Program struct {
	operations []Operation
}

// NewProgram function    Creates a new program
func NewProgram() *Program {
	return &Program{
		operations: []Operation{},
	}
}

// Operations method    Returns the operations of the program
func (p *Program) Operations() []Operation {
	return p.operations
}

// IsEmpty method    Returns true if the program is empty
func (p *Program) IsEmpty() bool {
	return len(p.operations) == 0
}

// Push method    Adds an operation to the program
func (p *Program) Push(op Operation) {
	p.operations = append(p.operations, op)
}
