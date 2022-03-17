package interpreter

import (
	"eslang/core"
	"fmt"
)

type StackValue struct {
	intValue   int64
	floatValue float64
	_type      core.Type
}

func NewStackValue() *StackValue {
	return &StackValue{
		_type: core.Nil,
	}
}

// TODO: Create a TypeAlias like Token alias
func (sv *StackValue) Type() core.Type {
	return sv._type
}

func (sv *StackValue) Value() (interface{}, error) {
	switch sv._type {
	case core.Int:
		return sv.Int(), nil
	case core.Float:
		return sv.Float(), nil
	}

	return nil, fmt.Errorf("exaustive type handiling for `%d`", sv._type)
}

func (sv *StackValue) Int() int64 {
	return sv.intValue
}

func (sv *StackValue) SetInt(value int64) *StackValue {
	sv.intValue = value
	sv._type = core.Int

	return sv
}

func (sv *StackValue) Float() float64 {
	return sv.floatValue
}

func (sv *StackValue) SetFloat(value float64) *StackValue {
	sv.floatValue = value
	sv._type = core.Float

	return sv
}

type Stack []*StackValue

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(op *StackValue) {
	*s = append(*s, op)
}

func (s *Stack) Pop() (*StackValue, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Stack.Pop()`, stack is empty.")
	}

	// Get the index of the top most element.
	index := len(*s) - 1
	// Index into the slice and obtain the element.
	value := (*s)[index]
	// Remove it from the stack by slicing it off.
	*s = (*s)[:index]

	return value, nil
}

func (s *Stack) PopTwo() (lhs *StackValue, rhs *StackValue, err error) {
	rhs, err = s.Pop()

	if err != nil {
		return nil, nil, err
	}

	lhs, err = s.Pop()

	if err != nil {
		return nil, nil, err
	}

	return lhs, rhs, nil
}

func executeProgram(program *core.Program, stack *Stack) error {
	for _, op := range *program {
		handler := REGISTERED_OPERATIONS[op.Type()]

		switch op.Type() {
		case core.OP_BLOCK:
			switch op.TokenStart().Token() {
			case core.TOKEN_DO:
				program := op.Value().Block().Current()

				if err := executeProgram(program, stack); err != nil {
					return FormatError(op, err)
				}
			case core.TOKEN_IF:
				sValue, err := stack.Pop()
				if err != nil {
					return FormatError(op, fmt.Errorf("error running block program"))
				}

				truthy := sValue.Int()

				// TODO: Add boolean type
				if sValue.Type() != core.Int {
					value, err := sValue.Value()
					if err != nil {
						return err
					}

					return FormatError(op, fmt.Errorf(
						"error testing the truthy of %s with type %d",
						value,
						sValue.Type(),
					),
					)
				}

				block := op.Value().Block()

				if truthy != 0 {
					program := block.Current()

					if err := executeProgram(program, stack); err != nil {
						return FormatError(op, err)
					}

					break
				}

				if !block.HasNext() {
					break
				}

				program := block.Next().Current()

				if err := executeProgram(program, stack); err != nil {
					return FormatError(op, err)
				}

			}

		default:
			if err := handler(stack, op); err != nil {
				return FormatError(op, err)
			}
		}
	}

	return nil
}

func SimulateProgram(program *core.Program) error {
	var stack Stack

	if err := executeProgram(program, &stack); err != nil {
		return err
	}

	return nil
}
