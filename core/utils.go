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
		case OP_BLOCK:
			lnum, cnum := op.TokenEnd().Position()
			eline := color.InBold(strconv.Itoa(lnum))
			ecol := color.InBold(strconv.Itoa(cnum))

			fmt.Printf("%s%s in lines [%s:%s:%s:%s]\n", spacing, token, line, col, eline, ecol)

			if err := PrintProgram(op.Value().Block().Current(), ident+1); err != nil {
				return err
			}

			block := op.Value().Block()

			if block.HasNext() {
				token := color.InYellow(TOKEN_ALIASES[TOKEN_ELSE])

				lnum, cnum := block.Next().TokenEnd().Position()
				eline := color.InBold(strconv.Itoa(lnum))
				ecol := color.InBold(strconv.Itoa(cnum))

				fmt.Printf("%s%s in lines [%s:%s:%s:%s]\n", spacing, token, line, col, eline, ecol)

				if err := PrintProgram(block.Next().Current(), ident+1); err != nil {
					return err
				}
			}

			endToken := color.InYellow(TOKEN_ALIASES[TOKEN_END])

			lnum, cnum = block.Last().TokenEnd().Position()
			line := color.InBold(strconv.Itoa(lnum))
			col := color.InBold(strconv.Itoa(cnum))

			fmt.Printf("%s%s in line %s:%s\n", spacing, endToken, line, col)

		default:
			var value string

			switch op.Type() {
			case OP_PUSH_STR:
				v := op.Value().Str()
				value = fmt.Sprintf("%v", v)
				value = fmt.Sprintf(" %s ", color.InCyan(value))
			case OP_PUSH_INT:
				v := op.Value().Int()
				value = fmt.Sprintf("%v", v)
				value = fmt.Sprintf(" %s ", color.InCyan(value))
			case OP_PUSH_FLOAT:
				v := op.Value().Int()
				value = fmt.Sprintf("%v", v)
				value = fmt.Sprintf(" %s ", color.InCyan(value))
			default:
				value = " "
			}

			fmt.Printf("%s%s%sin line %s:%s\n", spacing, token, value, line, col)
		}
	}

	return nil
}
