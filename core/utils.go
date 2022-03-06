package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
)

const LoggerTimeFormat = "2006/01/02 15:04:05"

func PrintProgram(program *Program, ident int) error {
	for _, op := range *program {
		lnum, cnum := op.TokenStart().Position()
		token := op.TokenStart().TokenAlias()
		spacing := strings.Repeat("\t", ident)

		token = color.InYellow(token)

		line := color.InBold(strconv.Itoa(lnum))
		col := color.InBold(strconv.Itoa(cnum))

		switch op.Type() {
		case OP_BLOCK, OP_IF:
			fmt.Printf("%s%s in line %s:%s\n", spacing, token, line, col)

			if err := PrintProgram(op.Value().(*Program), ident+1); err != nil {
				return err
			}

			block := op.(BlockOperation)

			if block.HasElseBlock() {
				token := color.InYellow(TOKEN_MAPPING[TOKEN_ELSE])

				fmt.Printf("%s%s in line %s:%s\n", spacing, token, line, col)

				if err := PrintProgram(block.ElseBlock(), ident+1); err != nil {
					return err
				}
			}

			endToken := color.InYellow(TOKEN_MAPPING[TOKEN_END])

			lnum, cnum := op.TokenEnd().Position()
			line := color.InBold(strconv.Itoa(lnum))
			col := color.InBold(strconv.Itoa(cnum))

			fmt.Printf("%s%s in line %s:%s\n", spacing, endToken, line, col)

		default:
			value := " "

			if op.Type() == OP_PUSH {
				v := op.Value().(int64)
				value = fmt.Sprintf("%v", v)
				value = fmt.Sprintf(" %s ", color.InCyan(value))
			}

			fmt.Printf("%s%s%sin line %s:%s\n", spacing, token, value, line, col)
		}
	}

	return nil
}
