package interpreter

import (
	"eslang/core"
	"fmt"
)

func FormatError(op *core.Operation, err error) error {
	line, col := op.TokenStart().Position()
	token := op.TokenStart().TokenAlias()

	return fmt.Errorf("%s - Token %s in line %d:%d", err.Error(), token, line, col)
}
