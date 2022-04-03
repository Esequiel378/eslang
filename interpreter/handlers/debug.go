package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

// OPDebug function    Prints the stack
func OPDebug(stack *s.Stack, _ ops.Operation) error {
	fmt.Println(stack.Content())

	return nil
}
