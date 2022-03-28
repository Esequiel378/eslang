package tokens

import (
	"eslang/core"
	ops "eslang/core/operations"
	"fmt"
	"regexp"
)

var IS_VALID_VARIABLE = regexp.MustCompile(`^[a-zA-Z]+$`)

var ALIASES_TO_TYPE = map[string]core.Type{
	"int":   core.Int,
	"float": core.Float,
	"str":   core.String,
}

// TokenVariable function    create and push a variable onto the stack
func TokenVariable(name string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_VALID_VARIABLE.MatchString(name) || core.RESERVED_WORDS[name] {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPVariable(name, nil, position)

	if variable, found := program.GetVariable(name); found {
		op = ops.NewOPVariable(name, variable.Value(), position)
	}

	err := program.Push(op)

	return true, err
}

// TokenSetVariableType function    set the last variable pushed to a specific type
func TokenSetVariableType(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	t, ok := ALIASES_TO_TYPE[token]

	if !ok {
		return false, nil
	}

	lastOP := program.LastOP()

	if lastOP == nil || lastOP.Type() != ops.OP_VARIABLE {
		return true, fmt.Errorf("using type `%s` out of context in line %d:%d", token, lnum, cnum)
	}

	position := ops.NewPosition(lnum, cnum, "")
	variable := lastOP.(*ops.OPVariable)

	if variable, found := program.GetVariable(variable.Name()); found {
		line, column := variable.Position().Ruler()
		file := variable.Position().File()
		name := variable.Name()

		return true, fmt.Errorf("variable `%s` already defined in %s:%d:%d", name, file, line, column)
	}

	var value ops.Operation

	switch t {
	case core.Int:
		value = ops.NewOPPushInt(0, position)
	case core.Float:
		value = ops.NewOPPushFloat(0.0, position)
	case core.String:
		value = ops.NewOPPushString("", position)
	default:
		return true, fmt.Errorf("unknown type `%s` in line %d:%d", token, lnum, cnum)
	}

	variable.SetValue(value)
	program.SetVariable(variable.Name(), variable)
	program.ChangeLastOP(variable)

	return true, nil
}

// TokenVariableWrite function    write a value into a variable
func TokenVariableWrite(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "." {
		return false, nil
	}

	lhs, rhs, err := program.PeekTwo()
	if err != nil {
		return false, err
	}

	variable := rhs.(*ops.OPVariable)

	variableValueType := variable.Value().Type()

	if lhs.Type() != variableValueType {
		return false, fmt.Errorf("cannot assign `%s` to `%s` in line %d:%d", lhs.Type(), variableValueType, lnum, cnum)
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPVariableWrite(position)

	err = program.Push(op)

	return true, err
}
