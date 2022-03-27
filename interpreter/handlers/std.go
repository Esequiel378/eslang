package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

func OPDump(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	fmt.Println(sValue.Value())

	return nil
}
