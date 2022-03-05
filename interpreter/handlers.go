package interpreter

import (
	"eslang/core"
	"fmt"
)

var REGISTERED_OPERATIONS = map[int]func(*core.Stack, core.Operation) error{
	core.OP_PUSH:  OPPush,
	core.OP_PLUS:  OPPlus,
	core.OP_MINUS: OPMinus,
	core.OP_EQUAL: OPEqual,
	core.OP_DUMP:  OPDump,
}

func OPPush(stack *core.Stack, op core.Operation) error {
	value, err := op.Value(stack)

	if err != nil {
		return err
	}

	stack.Push(value)

	return nil
}

func OPPlus(stack *core.Stack, op core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	lhsf, rhsf, keepFloat, err := normalizeNumbers(lhs, rhs)

	if err != nil {
		return err
	}

	if keepFloat {
		stack.Push(lhsf + rhsf)
	} else {
		stack.Push(int64(lhsf + rhsf))
	}

	return nil
}

func OPMinus(stack *core.Stack, op core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	lhsf, rhsf, keepFloat, err := normalizeNumbers(lhs, rhs)

	if err != nil {
		return err
	}

	if keepFloat {
		stack.Push(lhsf - rhsf)
	} else {
		stack.Push(int64(lhsf - rhsf))
	}

	return nil
}

func OPEqual(stack *core.Stack, op core.Operation) error {
	lhs, rhs, err := stack.PopLastTwo()

	if err != nil {
		return err
	}

	var bitSet int64

	// TODO: at some point this should be using bool type
	if lhs == rhs {
		bitSet = 1
	}

	stack.Push(bitSet)

	return nil
}

func OPDump(stack *core.Stack, op core.Operation) error {
	value, err := stack.Pop()

	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}

// func (b *IfBlockOperation) Value(stack *Stack) (interface{}, error) {
// 	value, err := stack.Pop()

// 	if err != nil {
// 		return nil, err
// 	}

// 	truthy, ok := value.(int64)

// 	if !ok {
// 		return nil, fmt.Errorf(
// 			"error testing the truthy of %s with type %s",
// 			value,
// 			reflect.TypeOf(value),
// 		)
// 	}

// 	if truthy != 0 {
// 		return b.block, nil
// 	}

// 	if isNil(b.elseBlock) {
// 		return nil, nil
// 	}

// 	return b.elseBlock, nil
// }
