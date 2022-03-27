package operations

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

// NewOPPushFloat function 
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
