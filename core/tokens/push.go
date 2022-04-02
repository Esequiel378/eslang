package tokens

import (
	ops "eslang/core/operations"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	IS_INT   = regexp.MustCompile(`^(\d{1,3})(_?\d{1,3})*$`)
	IS_FLOAT = regexp.MustCompile(`^\d+\.\d+$`)
	IS_STR   = regexp.MustCompile(`^".+"$`)
)

// TokenPushInt function    Push an 64 bytes int onto the stack
func TokenPushInt(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_INT.MatchString(token) {
		return false, nil
	}

	token = strings.ReplaceAll(token, "_", "")

	value, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return true, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushInt(value, position)

	err = program.Push(op)

	return true, err
}

// TokenPushFloat function    Push a 64 bytes float onto the stack
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

	err = program.Push(op)

	return true, err
}

// TokenPushStr function    Push a string onto the stack
func TokenPushStr(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if !IS_STR.MatchString(token) {
		return false, nil
	}

	value := token[1 : len(token)-1]
	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushString(value, position)

	err := program.Push(op)

	return true, err
}

// TokenPushBool function    Push a boolean onto the stack
func TokenPushBool(token string, lnum, cnum int, program *ops.Program) (bool, error) {
	if token != "true" && token != "false" {
		return false, nil
	}

	value := token == "true"
	position := ops.NewPosition(lnum, cnum, "")
	op := ops.NewOPPushBool(value, position)

	err := program.Push(op)

	return true, err
}
