package utils

import (
	ops "eslang/core/operations"
	tkns "eslang/core/tokens"
	"fmt"
	"os"
	"strings"

	"github.com/anmitsu/go-shlex"
)

func LoadProgramFromFile(program *ops.Program, filename string) error {
	rawLines, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(rawLines), "\n")

	for lnum, line := range lines {
		line = strings.Split(line, "//")[0]

		if len(line) == 0 || line == "" {
			continue
		}

		tokens, err := shlex.Split(line, false)
		if err != nil {
			return err
		}

		for _, token := range tokens {
			token = strings.Trim(token, " ")

			if len(token) == 0 {
				continue
			}

			found := false
			cnum := strings.Index(line, token)

			for _, tokenHandler := range tkns.REGISTERED_TOKENS {
				// TODO: send filename to toke handler
				ok, err := tokenHandler(token, lnum+1, cnum+1, program)
				if err != nil {
					return err
				}

				if !ok {
					continue
				}

				found = true
				break
			}

			if !found {
				// TODO: improve error mesage using colors and stuff
				return fmt.Errorf(
					"token error in %d:%d - '%s' is not a valid token",
					lnum+1, cnum+1,
					token,
				)
			}
		}
	}

	lastOP := program.LastOP()

	if lastOP.Type().IsBlock() {
		block := lastOP.(ops.OperationBlock)

		if !block.IsClosed() {
			line, column := block.Position().Ruler()

			// Check if the block has liked blocks and update the line and column values
			if b, ok := block.(ops.OperationLinkedBlocks); ok {
				fmt.Println("Here!")
				block := b.LastBlock().(ops.OperationBlock)
				line, column = block.Position().Ruler()
			}

			return fmt.Errorf("unclosed block at %s:%d:%d", filename, line, column)
		}
	}

	return nil
}
