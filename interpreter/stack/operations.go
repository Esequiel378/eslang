package stack

import (
	"eslang/core"
	"fmt"
	"math"
)

// ======================
// ARITHMETIC OPERATIONS
// ======================

// AddValues function    adds two values together
func AddValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not add values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) + rhs.Value().(int64)), nil
	case core.Float:
		return NewStackValueFloat(lhs.Value().(float64) + rhs.Value().(float64)), nil
	case core.String:
		return NewStackValueString(lhs.Value().(string) + rhs.Value().(string)), nil
	}

	return nil, fmt.Errorf("can not add values of type %s", lhs.Type())
}

// SubtractValues function    subtracts two values
func SubtractValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not subtract values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) - rhs.Value().(int64)), nil
	case core.Float:
		return NewStackValueFloat(lhs.Value().(float64) - rhs.Value().(float64)), nil
	}

	return nil, fmt.Errorf("can not subtract values of type %s", lhs.Type())
}

// MultiplyValues function    multiplies two values
func MultiplyValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not multiply values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) * rhs.Value().(int64)), nil
	case core.Float:
		return NewStackValueFloat(lhs.Value().(float64) * rhs.Value().(float64)), nil
	}

	return nil, fmt.Errorf("can not multiply values of type %s", lhs.Type())
}

// DivideValues function  Óòß  divides two values
func DivideValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not divide values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) / rhs.Value().(int64)), nil
	case core.Float:
		return NewStackValueFloat(lhs.Value().(float64) / rhs.Value().(float64)), nil
	}

	return nil, fmt.Errorf("can not divide values of type %s", lhs.Type())
}

// ModuloValues function    returns the remainder of the divition between two values
func ModuloValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not modulo values of different types")
	}

	switch lhs.Type() {
	case core.Int:
		return NewStackValueInt(lhs.Value().(int64) % rhs.Value().(int64)), nil
	case core.Float:
		return NewStackValueFloat(
			math.Mod(
				lhs.Value().(float64),
				rhs.Value().(float64),
			),
		), nil
	}

	return nil, fmt.Errorf("can not modulo values of type %s", lhs.Type())
}

// ======================
// RELATIONAL OPERATIONS
// ======================

// EqualValues function    compares if two values are equal
func EqualValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	var truthty int64 = 0

	if lhs.Value() == rhs.Value() {
		truthty = 1
	}

	return NewStackValueInt(truthty), nil
}

// NotEqualValues function    compares if two values are not equal
func NotEqualValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	var truthty int64 = 0

	if lhs.Value() != rhs.Value() {
		truthty = 1
	}

	return NewStackValueInt(truthty), nil
}

// LessThanValues function    compares if one value is less than another
func LessThanValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	var truthty int64 = 0

	switch lhs.Type() {
	case core.Int:
		if lhs.Value().(int64) < rhs.Value().(int64) {
			truthty = 1
		}
	case core.Float:
		if lhs.Value().(float64) < rhs.Value().(float64) {
			truthty = 1
		}
	case core.String:
		if lhs.Value().(string) < rhs.Value().(string) {
			truthty = 1
		}
	default:
		return nil, fmt.Errorf("can not compare values of type %s", lhs.Type())
	}

	return NewStackValueInt(truthty), nil
}

// GreaterThanValues function    compares if one value is greater than another
func GreaterThanValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	var truthty int64 = 0

	switch lhs.Type() {
	case core.Int:
		if lhs.Value().(int64) > rhs.Value().(int64) {
			truthty = 1
		}
	case core.Float:
		if lhs.Value().(float64) > rhs.Value().(float64) {
			truthty = 1
		}
	case core.String:
		if lhs.Value().(string) > rhs.Value().(string) {
			truthty = 1
		}
	default:
		return nil, fmt.Errorf("can not compare values of type %s", lhs.Type())
	}

	return NewStackValueInt(truthty), nil
}

// LessThanOrEqualValues function    compares if one value is less than or equal to another
func LessThanOrEqualValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	var truthty int64 = 0

	switch lhs.Type() {
	case core.Int:
		if lhs.Value().(int64) <= rhs.Value().(int64) {
			truthty = 1
		}
	case core.Float:
		if lhs.Value().(float64) <= rhs.Value().(float64) {
			truthty = 1
		}
	case core.String:
		if lhs.Value().(string) <= rhs.Value().(string) {
			truthty = 1
		}
	default:
		return nil, fmt.Errorf("can not compare values of type %s", lhs.Type())
	}

	return NewStackValueInt(truthty), nil
}

// GreaterThanOrEqualValues function    compares if one value is greater than or equal to another
func GreaterThanOrEqualValues(lhs StackValue, rhs StackValue) (StackValue, error) {
	if lhs.Type() != rhs.Type() {
		return nil, fmt.Errorf("can not compare values of different types")
	}

	var truthty int64 = 0

	switch lhs.Type() {
	case core.Int:
		if lhs.Value().(int64) >= rhs.Value().(int64) {
			truthty = 1
		}
	case core.Float:
		if lhs.Value().(float64) >= rhs.Value().(float64) {
			truthty = 1
		}
	case core.String:
		if lhs.Value().(string) >= rhs.Value().(string) {
			truthty = 1
		}
	default:
		return nil, fmt.Errorf("can not compare values of type %s", lhs.Type())
	}

	return NewStackValueInt(truthty), nil
}
