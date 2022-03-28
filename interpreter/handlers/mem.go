package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
)

func OPVariable(stack *s.Stack, _op ops.Operation) error {
	op := _op.(*ops.OPVariable)

	name := op.Name()

	variable, found := stack.GetVariable(name)

	if !found {
		variable = s.NewStackValueVar(name, nil)
		stack.SetVariable(name, variable)
	}

	stack.Push(variable)

	return nil
}
