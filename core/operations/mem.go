package operations

// OPVariable struct    Hold a value in memory
type OPVariable struct {
	value    Operation
	position Position
	name     string
}

// NewOPVariable function    Create a new variable
func NewOPVariable(name string, value Operation, position Position) *OPVariable {
	return &OPVariable{
		name:     name,
		value:    value,
		position: position,
	}
}

// Value method    Get the value of the variable
func (o OPVariable) Value() Operation {
	return o.value
}

// SetValue method    Set the value of the variable
func (o *OPVariable) SetValue(value Operation) {
	o.value = value
}

// Type method    Get the type of the operation
func (o OPVariable) Type() OPType {
	return OP_VARIABLE
}

// Postition method    Get the position of the operation
func (o OPVariable) Position() Position {
	return o.position
}

// Name method    Get the name of the variable
func (o OPVariable) Name() string {
	return o.name
}
