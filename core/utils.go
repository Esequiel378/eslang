package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
)

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
			lnum, cnum := op.TokenEnd().Position()
			eline := color.InBold(strconv.Itoa(lnum))
			ecol := color.InBold(strconv.Itoa(cnum))

			fmt.Printf("%s%s in lines [%s:%s:%s:%s]\n", spacing, token, line, col, eline, ecol)

			if err := PrintProgram(op.Value().(*Program), ident+1); err != nil {
				return err
			}

			block := op.(BlockOperation)

			if block.HasRefBlock() {
				token := color.InYellow(TOKEN_MAPPING[TOKEN_ELSE])

				lnum, cnum := block.RefBlock().TokenEnd().Position()
				eline := color.InBold(strconv.Itoa(lnum))
				ecol := color.InBold(strconv.Itoa(cnum))

				fmt.Printf("%s%s in lines [%s:%s:%s:%s]\n", spacing, token, line, col, eline, ecol)

				if err := PrintProgram(block.RefBlock().Block(), ident+1); err != nil {
					return err
				}
			}

			endToken := color.InYellow(TOKEN_MAPPING[TOKEN_END])

			lnum, cnum = block.LastBlock().TokenEnd().Position()
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
