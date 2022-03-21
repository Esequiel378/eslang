package interpreter

import (
	"eslang/core"
	"fmt"
)

type StackValue struct {
	intValue   int64
	floatValue float64
	strValue   string
	name       string
	_type      core.Type
}

func NewStackValue() *StackValue {
	return &StackValue{
		_type: core.Nil,
	}
}

func (sv *StackValue) Name() string {
	return sv.name
}

func (sv *StackValue) SetName(name string) *StackValue {
	sv.name = name

	return sv
}

func (sv *StackValue) Type() core.Type {
	return sv._type
}

func (sv *StackValue) SetType(t core.Type) *StackValue {
	sv._type = t

	return sv
}

func (sv *StackValue) TypeAlias() string {
	return core.TYPE_ALIASES[sv._type]
}

func (sv *StackValue) Value() (interface{}, error) {
	switch sv._type {
	case core.Int:
		return sv.Int(), nil
	case core.Float:
		return sv.Float(), nil
	case core.Str:
		return sv.Str(), nil
	}

	return nil, fmt.Errorf("exaustive type handiling for `%d`", sv._type)
}

func (sv StackValue) IsNumber() bool {
	return sv._type == core.Int || sv._type == core.Float
}

func (sv *StackValue) Str() string {
	return sv.strValue
}

func (sv *StackValue) SetStr(str string) *StackValue {
	sv.strValue = str
	sv._type = core.Str

	return sv
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

type Stack struct {
	stack     []*StackValue
	variables map[string]*StackValue
}

func NewStack() Stack {
	return Stack{
		stack:     []*StackValue{},
		variables: make(map[string]*StackValue),
	}
}

func (s *Stack) SetVariable(name string, sValue *StackValue) {
	s.variables[name] = sValue
}

func (s *Stack) GetVariable(name string) (*StackValue, bool) {
	v, found := s.variables[name]

	return v, found
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Push(op *StackValue) {
	s.stack = append(s.stack, op)
}

func (s *Stack) Pop() (*StackValue, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Stack.Pop()`, stack is empty")
	}

	// Get the index of the top most element.
	index := len(s.stack) - 1
	// Index into the slice and obtain the element.
	value := (s.stack)[index]
	// Remove it from the stack by slicing it off.
	s.stack = (s.stack)[:index]

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
	for _, op := range program.Operations() {
		switch op.Type() {
		case core.OP_BLOCK:
			handler, ok := REGISTERED_OP_BLOCK[op.TokenStart().Token()]
			if !ok {
				return FormatError(
					op,
					fmt.Errorf("exaustive block operation handiling for `%s`", op.TypeAlias()),
				)
			}

			program, err := handler(stack, op)
			if err != nil {
				return FormatError(op, err)
			}

			if program == nil {
				break
			}

			if err := executeProgram(program, stack); err != nil {
				return FormatError(op, err)
			}

		default:
			handler, ok := REGISTERED_OPERATIONS[op.Type()]

			if !ok {
				return FormatError(
					op,
					fmt.Errorf("exaustive operation handiling for `%s`", op.TypeAlias()),
				)
			}

			if err := handler(stack, op); err != nil {
				return err
			}
		}
	}

	return nil
}

func SimulateProgram(program *core.Program) error {
	stack := NewStack()

	if err := executeProgram(program, &stack); err != nil {
		return err
	}

	return nil
}
