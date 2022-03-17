package interpreter

import (
	"eslang/core"
	"fmt"
	"reflect"
)

var floatType = reflect.TypeOf(float64(0))

func getFloat64(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)

	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}

	fv := v.Convert(floatType)
	return fv.Float(), nil
}

func normalizeNumbers(lhs, rhs *StackValue) {
	parseFloat := lhs.Type() == core.Float || rhs.Type() == core.Float

	if parseFloat {
		lhs.SetFloat(float64(lhs.Int()))
		rhs.SetFloat(float64(rhs.Int()))
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}

	return false
}

func FormatError(op *core.Operation, err error) error {
	line, col := op.TokenStart().Position()
	token := op.TokenStart().TokenAlias()

	return fmt.Errorf("error `%s` with Token: %s in line %d:%d", err.Error(), token, line, col)
}
