package interpreter

import (
	"eslang/core"
	s "eslang/interpreter/stack"
	"fmt"
)

type (
	OPHandler      func(*s.Stack, *core.Operation) error
	OPBlockHandler func(*s.Stack, *core.Operation) (*core.Program, error)
)

var REGISTERED_OP_BLOCK = map[core.TokenType]OPBlockHandler{
	core.TOKEN_DO:    OPBlockDo,
	core.TOKEN_IF:    OPBlockIf,
	core.TOKEN_WHILE: OPBlockIf,
}

func OPBlockDo(_ *s.Stack, op *core.Operation) (*core.Program, error) {
	program := op.Value().Block().Program()

	return program, nil
}

func OPBlockIf(stack *s.Stack, op *core.Operation) (*core.Program, error) {
	sValue, err := stack.Peek()
	if err != nil {
		return nil, FormatError(op, err)
	}

	truthy, err := sValue.TestTruthy()
	if err != nil {
		return nil, FormatError(op, err)
	}

	block := op.Value().Block()

	// If block
	if truthy {
		program := block.Program()
		return program, nil
	}

	// Else block
	if block.HasNext() {
		return block.Next().Program(), nil
	}

	// End block
	return nil, nil
}

var REGISTERED_OPERATIONS = map[core.OperationType]OPHandler{
	core.OP_DUMP:       OPDump,
	core.OP_DUP:        OPDup,
	core.OP_MOP:        OPMop,
	core.OP_PUSH_FLOAT: OPPushFloat,
	core.OP_PUSH_INT:   OPPushInt,
	core.OP_PUSH_STR:   OPPushStr,
	// core.OP_VAR:        OPVar,
	// core.OP_VAR_WRITE:  OPVarWrite,
}

func OPPushFloat(stack *s.Stack, op *core.Operation) error {
	sValue := s.NewStackValueFloat(op.Value().Float())
	stack.Push(sValue)

	return nil
}

func OPPushInt(stack *s.Stack, op *core.Operation) error {
	sValue := s.NewStackValueInt(op.Value().Int())
	stack.Push(sValue)

	return nil
}

func OPPushStr(stack *s.Stack, op *core.Operation) error {
	sValue := s.NewStackValueString(op.Value().Str())
	stack.Push(sValue)

	return nil
}

func OPDump(stack *s.Stack, _ *core.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	fmt.Println(sValue.Value())

	return nil
}

func OPDup(stack *s.Stack, _ *core.Operation) error {
	sValue, err := stack.Pop()
	if err != nil {
		return err
	}

	stack.Push(sValue)
	stack.Push(sValue)

	return nil
}

var REGISTERED_MOPS = map[core.TokenType]func(*s.Stack, *core.Operation) error{
	core.TOKEN_EQUAL: OPEqual,
	core.TOKEN_MINUS: OPMinus,
	core.TOKEN_PLUS:  OPPlus,
}

func OPMop(stack *s.Stack, op *core.Operation) error {
	handler, ok := REGISTERED_MOPS[op.TokenStart().Token()]

	if !ok {
		return fmt.Errorf("exaustive MOPs handiling. %s not found", op.TokenEnd().TokenAlias())
	}

	if err := handler(stack, op); err != nil {
		return err
	}

	return nil
}

func OPPlus(stack *s.Stack, _ *core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	sValue, err := s.AddValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(sValue)

	return nil
}

func OPMinus(stack *s.Stack, _ *core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	sValue, err := s.SubtractValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(sValue)

	return nil
}

func OPEqual(stack *s.Stack, _ *core.Operation) error {
	lhs, rhs, err := stack.PopTwo()
	if err != nil {
		return err
	}

	sValue, err := s.CompareEqualValues(lhs, rhs)
	if err != nil {
		return err
	}

	stack.Push(sValue)

	return nil
}

// func OPVar(stack *s.Stack, op *core.Operation) error {
// 	opValue := op.Value()

// 	name := op.Value().Name()
// 	variable, found := stack.GetVariable(name)

// 	if !found {
// 		variable = NewStackValue().SetName(name).SetType(opValue.Type())
// 		stack.SetVariable(name, variable)
// 	}

// 	stack.Push(variable)

// 	return nil
// }

// func OPVarWrite(stack *s.Stack, _ *core.Operation) error {
// 	lhs, rhs, err := stack.PopTwo()
// 	if err != nil {
// 		return err
// 	}

// 	sValue := NewStackValue().SetName(rhs.Name())

// 	if rhs.Name() == "" && lhs.Name() == "" {
// 		return fmt.Errorf("`write` operation can only be used with variables")
// 	}

// 	if rhs.Name() == "" {
// 		return fmt.Errorf("error writing to variable, invalid parameters order")
// 	}

// 	switch rhs.Type() {
// 	case core.Int:
// 		if lhs.Type() != core.Int {
// 			return fmt.Errorf("can not assign an `%s` value to a `%s` variable", lhs.TypeAlias(), rhs.TypeAlias())
// 		}

// 		sValue.SetInt(lhs.Int())
// 	case core.Float:
// 		if lhs.Type() != core.Float {
// 			return fmt.Errorf("can not assign an `%s` value to a `%s` variable", lhs.TypeAlias(), rhs.TypeAlias())
// 		}

// 		sValue.SetFloat(lhs.Float())
// 	case core.String:
// 		if lhs.Type() != core.String {
// 			return fmt.Errorf("can not assign an `%s` value to a `%s` variable", lhs.TypeAlias(), rhs.TypeAlias())
// 		}

// 		sValue.SetStr(lhs.Str())
// 	}

// 	stack.SetVariable(rhs.Name(), sValue)

// 	stack.Push(sValue)

// 	return nil
// }
