package interpreter

import (
	"eslang/core"
	"fmt"
)

func FormatError(op core.Operation, err error) error {
	line, col := op.Position().Ruler()
	file := op.Position().File()

	return fmt.Errorf("%s:%d:%d - %s in %s", file, line, col, err.Error(), op.Type())
}
