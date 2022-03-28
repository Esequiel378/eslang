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

	if !found {
		variable = s.NewStackValueVariable(name, nil)
		stack.SetVariable(name, variable)
	}

	stack.Push(variable)

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
	stack.Push(sValue)

	return nil
}
