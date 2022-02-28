package interpreter

import (
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

func normalizeNumbers(lhs, rhs interface{}) (lhsf, rhsf float64, err error) {
	lhsf, err = getFloat64(lhs)

	if err != nil {
		return 0, 0, err
	}

	rhsf, err = getFloat64(rhs)

	return lhsf, rhsf, err
}
