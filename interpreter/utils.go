package interpreter

import (
	"eslang/core"
	"fmt"
)

func normalizeNumbers(lhs, rhs *StackValue) {
	parseFloat := lhs.Type() == core.Float || rhs.Type() == core.Float

	if parseFloat {
		lhs.SetFloat(float64(lhs.Int()))
		rhs.SetFloat(float64(rhs.Int()))
	}
}

func FormatError(op *core.Operation, err error) error {
	line, col := op.TokenStart().Position()
	token := op.TokenStart().TokenAlias()

	return fmt.Errorf("error: %s - Token %s in line %d:%d", err.Error(), token, line, col)
}
