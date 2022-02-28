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

func normalizeNumbers(lhs, rhs interface{}) (lhsf, rhsf float64, keepFloat bool, err error) {
	lhsf, err = getFloat64(lhs)

	if err != nil {
		return 0, 0, false, err
	}

	rhsf, err = getFloat64(rhs)

	if err != nil {
		return lhsf, 0, false, err
	}

	keepFloat = false

	numbers := []interface{}{lhs, rhs}

	for _, number := range numbers {
		if _, ok := number.(float64); ok {
			keepFloat = true
			break
		}
	}

	return lhsf, rhsf, keepFloat, nil
}
