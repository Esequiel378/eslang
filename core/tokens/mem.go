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

func TokenVariable(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_VALID_VARIABLE.MatchString(token) || core.RESERVED_WORDS[token] {
		return false, nil
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPVariable(token, nil, position)

	err := program.Push(op)

	return true, err
}

func TokenSetVariableType(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	t, ok := ALIASES_TO_TYPE[token]

	if !ok {
		return false, nil
	}

	lastOP := program.LastOP()

	if lastOP == nil || lastOP.Type() != ops.OP_VARIABLE {
		return true, fmt.Errorf("using type `%s` out of context at line %d:%d", token, lnum+1, cnum+1)
	}

	position := ops.NewPosition(lnum, cnum, "")
	variable := lastOP.(*ops.OPVariable)

	switch t {
	case core.Int:
		value := ops.NewOPPushInt(0, position)
		variable.SetValue(value)
	case core.Float:
		value := ops.NewOPPushFloat(0.0, position)
		variable.SetValue(value)
	case core.String:
		value := ops.NewOPPushString("", position)
		variable.SetValue(value)
	default:
		return true, fmt.Errorf("unknown type `%s` at line %d:%d", token, lnum, cnum)
	}

	return true, nil
}
