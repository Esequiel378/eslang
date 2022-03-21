package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/anmitsu/go-shlex"
)

type Program []*Operation

func (p *Program) IsEmpty() bool {
	return len(*p) == 0
}

// Push method    add a new operation to the program
func (p *Program) Push(op *Operation) {
	if !p.IsEmpty() {
		lastOp := p.Last()
		// Get the last block within the operation
		block := lastOp.Value().Block().Last()

		// Push the op to the last nested block if its open
		if lastOp.Type() == OP_BLOCK && block.IsOpen() {
			block.Current().Push(op)
			return
		}
	}

	*p = append(*p, op)
}

func (p *Program) CloseLastBlock(line, col int) error {
	if p.IsEmpty() {
		return fmt.Errorf("no open block to close")
	}

	lastOp := p.Last()

	if lastOp == nil || lastOp.Type() != OP_BLOCK {
		return fmt.Errorf("no open block to close")
	}

	// Get the last block within the operation
	block := lastOp.Value().Block().Last()

	block.TokenEnd().SetPostition(line, col)
	lastOp.TokenEnd().SetPostition(line, col)

	return nil
}

// Last method    Returns the last operation added to the program
// If there is not operation, it will return nil
func (p *Program) Last() *Operation {
	if p.IsEmpty() {
		return nil
	}

	return (*p)[len(*p)-1]
}

// TODO: Make parseLines an util not a method of Program
func (p *Program) parseLines(lines []string) error {
	for lnum, line := range lines {
		line = strings.Split(line, "//")[0]
		line = strings.Trim(line, " ")

		if len(line) == 0 {
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

			for _, tokenHandler := range REGISTERED_TOKENS {
				ok, err := tokenHandler(token, line, lnum, p)
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
				cnum := strings.Index(line, token)

				return fmt.Errorf(
					"Token error in %d:%d - '%s' is not a valid token",
					lnum+1, cnum+1,
					token,
				)
			}
		}
	}

	lastOp := p.Last()
	// Get the last block within the operation
	block := lastOp.Value().Block().Last()

	// Check for un-closed blocks
	if lastOp.Type() == OP_BLOCK && block.IsOpen() {
		tokenStart := block.TokenStart().TokenAlias()
		tokenEnd := block.TokenEnd().TokenAlias()
		lnum, cnum := block.TokenStart().Position()

		return fmt.Errorf(
			"%s missing %s closing token in line %d:%d",
			color.InYellow(tokenStart),
			color.InYellow(tokenEnd),
			lnum, cnum,
		)
	}

	return nil
}

func (p *Program) LoadFromFile(filename string) error {
	rawLines, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(rawLines), "\n")

	if err := p.parseLines(lines); err != nil {
		log.Fatal(err)
	}

	// TODO: check for a variable definitions instead
	for _, op := range *p {
		if op.Type() == OP_MEM && op.Value().Type() == Nil {
			line, col := op.TokenStart().Position()

			return fmt.Errorf(
				"missing variable type assigment in line %d:%d",
				line, col,
			)
		}
	}

	return nil
}
