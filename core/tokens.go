package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var REGISTERED_TOKENS = map[string]func(string) (*Operation, error){
	"PUSH":  TokenPush,
	"PLUS":  TokenPlus,
	"MINUS": TokenMinus,
	"EQUAL": TokenEqual,
	"DUMP":  TokenDump,
}

var IS_DIGIT = regexp.MustCompile(`^[1-9]\d*(\.\d+)?$`)

func tokenPushFloat64(token string) (*Operation, error) {
	value, err := strconv.ParseFloat(token, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to float: %s", token, err.Error())
	}

	operation := Operation{
		Type:  OP_PUSH,
		Value: value,
	}

	return &operation, nil
}

func tokenPushInt64(token string) (*Operation, error) {
	value, err := strconv.ParseInt(token, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing token '%s' to int: %s", token, err.Error())
	}

	operation := Operation{
		Type:  OP_PUSH,
		Value: value,
	}

	return &operation, nil
}

func TokenPush(token string) (*Operation, error) {
	if !IS_DIGIT.MatchString(token) {
		return nil, fmt.Errorf("Invalid token")
	}

	if strings.Contains(token, ".") {
		return tokenPushFloat64(token)
	}

	return tokenPushInt64(token)
}

func TokenPlus(token string) (*Operation, error) {
	if token != "+" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_PLUS,
	}

	return &operation, nil
}

func TokenMinus(token string) (*Operation, error) {
	if token != "-" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_MINUS,
	}

	return &operation, nil
}

func TokenEqual(token string) (*Operation, error) {
	if token != "=" {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_EQUAL,
	}

	return &operation, nil
}

func TokenDump(token string) (*Operation, error) {
	if token != "." {
		return nil, fmt.Errorf("Invalid token")
	}

	operation := Operation{
		Type: OP_DUMP,
	}

	return &operation, nil
}
