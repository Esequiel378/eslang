package handlers

import (
	ops "eslang/core/operations"
	s "eslang/interpreter/stack"
	"fmt"
)

// OPPushFloat function    Pushes a float value onto the stack.
func OPPushFloat(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushFloat)

	sValue := s.NewStackValueFloat(op.Value())
	stack.Push(sValue)

	return nil
}

// OPPushInt function    Pushes an integer value onto the stack.
func OPPushInt(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushInt)

	sValue := s.NewStackValueInt(op.Value())
	stack.Push(sValue)

	return nil
}

// OPPushStr function    Pushes a string value onto the stack.
func OPPushStr(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushString)

	sValue := s.NewStackValueString(op.Value())
	stack.Push(sValue)

	return nil
}

// OPPushBool function    Pushes a boolean value onto the stack.
func OPPushBool(stack *s.Stack, _op ops.Operation) error {
	op := _op.(ops.OPPushBool)

	sValue := s.NewStackValueBool(op.Value())
	stack.Push(sValue)

	return nil
}

// OPDrop function    drop the top of the stack
// Ex: ( a -- )
func OPDrop(stack *s.Stack, _ ops.Operation) error {
	_, err := stack.Pop()

	return err
}

// OPDump function    dump the top of the stack
// Ex: ( a -- )
func OPDump(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	fmt.Println(sValue.Value())

	return nil
}

// OPDup function    duplicate the top of the stack
// Ex: ( a -- a a )
func OPDup(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.Peek()
	if err != nil {
		return err
	}

	stack.Push(sValue)

	return nil
}

// OPOver function    duplicate the second-to-top of the stack
// Ex: ( a b -- a b a )
func OPOver(stack *s.Stack, _ ops.Operation) error {
	sValue, err := stack.PeekAt(-2)
	if err != nil {
		return err
	}

	stack.Push(sValue)

	return nil
}

// OPSwap function    swap the top two items on the stack
// Ex: ( a b -- b a )
func OPSwap(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(rhs)
	stack.Push(lhs)

	return nil
}

// OPRot function    rotates the top three items on the stack
// Ex: ( a b c -- c a b )
func OPRot(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	rotValue, err := stack.Pop()
	if err != nil {
		return err
	}

	stack.Push(lhs)
	stack.Push(rhs)

	stack.Push(rotValue)

	return nil
}

// OPORot function    rotates the top three items on the stack in the opposite direction
// Ex: ( a b c -- c a b )
func OPORot(stack *s.Stack, _ ops.Operation) error {
	rotValue, err := stack.Pop()
	if err != nil {
		return err
	}

	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(rotValue)

	stack.Push(lhs)
	stack.Push(rhs)

	return nil
}

// OPTuck function    duplicates the top of the stack and places it below the second-to-top
// Ex: ( a b -- a b a )
func OPTuck(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(rhs)
	stack.Push(lhs)
	stack.Push(rhs)

	return nil
}

// OPTwoDup function    duplicates the top two items on the stack
// Ex: ( a b -- a b a b )
func OPTwoDup(stack *s.Stack, _ ops.Operation) error {
	lhs, rhs, err := stack.PeekTwo()
	if err != nil {
		return err
	}

	stack.Push(lhs)
	stack.Push(rhs)

	return nil
}

// OpNip function    Removes the second-to-top item from the stack
// Ex: ( a b -- b )
func OPNip(stack *s.Stack, _ ops.Operation) error {
	lhs, _, er := stack.PopTwo()
	if er != nil {
		return er
	}

	stack.Push(lhs)

	return nil
}

// OPTwoDrop function    Removes the top two items from the stack
// Ex: ( a b -- )
func OPTwoDrop(stack *s.Stack, _ ops.Operation) error {
	_, _, err := stack.PopTwo()

	return err
}

// OPTwoOver function    Copies the second-to-top two items to the top of the stack
// Ex: ( a b c d -- a b c d a b )
func OPTwoOver(stack *s.Stack, _ ops.Operation) error {
	lhs_a, err := stack.PeekAt(-4)
	if err != nil {
		return err
	}
	stack.Push(lhs_a)

	rhs_b, err := stack.PeekAt(-4)
	if err != nil {
		return err
	}

	stack.Push(rhs_b)

	return nil
}

// OPTwoSwap function    Swaps the second-to-top two items on the stack
// Ex: ( a b c d -- c d a b )
func OPTwoSwap(stack *s.Stack, _ ops.Operation) error {
	lhs_a, rhs_b, err := stack.PopTwo()
	if err != nil {
		return err
	}

	lhs_c, rhs_d, err := stack.PopTwo()
	if err != nil {
		return err
	}

	stack.Push(lhs_c)
	stack.Push(rhs_d)

	stack.Push(lhs_a)
	stack.Push(rhs_b)

	return nil
}
