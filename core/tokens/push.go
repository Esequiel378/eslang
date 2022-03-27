package tokens

import (
	ops "eslang/core/operations"
	"fmt"
	"regexp"
	"strconv"
)

var (
	IS_INT   = regexp.MustCompile(`^\d+$`)
	IS_FLOAT = regexp.MustCompile(`^\d+\.\d+$`)
	IS_STR   = regexp.MustCompile(`^".+"$`)
)

// TokenPushInt function  
func TokenPushInt(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_INT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushInt(value, position)

	program.Push(op)

	return true, nil
}

// TokenPushFloat function    pushes a float onto the stack
func TokenPushFloat(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_FLOAT.MatchString(token) {
		return false, nil
	}

	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushFloat(value, position)

	program.Push(op)

	return true, nil
}

// TokenPushStr function    pushes a string onto the stack
func TokenPushStr(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_STR.MatchString(token) {
		return false, nil
	}

	value := token[1 : len(token)-1]
	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushString(value, position)

	program.Push(op)

	return true, nil
}
