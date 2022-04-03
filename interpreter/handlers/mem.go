package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

// OPVariable function    push a variable onto the stack if exists or create a new one
func OPVariable(stack *s.Stack, _op ops.Operation) error {
	op := _op.(*ops.OPVariable)

	name := op.Name()

	variable, found := stack.GetVariable(name)

	if found {
		stack.Push(variable)
		return nil
	}

	variableValue := op.Value()
	var value s.StackValue

	switch op.Value().Type() {
	case ops.OP_PUSH_INT:
		v := variableValue.(ops.OPPushInt)
		value = s.NewStackValueInt(v.Value())
	case ops.OP_PUSH_FLOAT:
		v := variableValue.(ops.OPPushFloat)
		value = s.NewStackValueFloat(v.Value())
	case ops.OP_PUSH_STRING:
		v := variableValue.(ops.OPPushString)
		value = s.NewStackValueString(v.Value())
	}

	variable = s.NewStackValueVariable(name, value)
	stack.SetVariable(name, variable)

	return nil
}

// OPVariableWrite function    write a value (lhs) to a variable (rhs)
func OPVariableWrite(stack *s.Stack, _ ops.Operation) error {
	lhs, _rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	rhs, ok := _rhs.(s.StackValueVar)
	if !ok {
		return fmt.Errorf("cannot write to non-variable")
	}

	sValue := s.NewStackValueVariable(rhs.Name(), lhs)

	stack.SetVariable(sValue.Name(), sValue)

	return nil
}
