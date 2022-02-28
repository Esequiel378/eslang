package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Program []Operation

func (p *Program) IsEmpty() bool {
	return len(*p) == 0
}

func (p *Program) Push(operation Operation) {
	*p = append(*p, operation)
}

func (p *Program) Pop() (Operation, error) {
	if p.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Program.Pop()`, program is empty.")
	}

	// Get the index of the top most element.
	index := len(*p) - 1
	// Index into the slice and obtain the element.
	operation := (*p)[index]
	// Remove it from the stack by slicing it off.
	*p = (*p)[:index]

	return operation, nil
}

type BlockStack []*BlockOperation

func (ps *BlockStack) IsEmpty() bool {
	return len(*ps) == 0
}

func (ps *BlockStack) Push(value *BlockOperation) {
	*ps = append(*ps, value)
}

func (ps *BlockStack) Pop() (*BlockOperation, error) {
	if ps.IsEmpty() {
		return nil, fmt.Errorf("can not perform `ProgramStack.Pop()`, stack is empty.")
	}

	// Get the index of the top most element.
	index := len(*ps) - 1
	// Index into the slice and obtain the element.
	value := (*ps)[index]
	// Remove it from the stack by slicing it off.
	*ps = (*ps)[:index]

	return value, nil
}

func getProgram(lines []string) (*Program, error) {
	var program Program

	var blocks BlockStack

	for lineNumber, line := range lines {
		line = strings.Trim(line, " ")

		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")

		for colNumber, token := range tokens {
			found := false

			for tokenType, tokenHandler := range REGISTERED_MISC_TOKENS {
				operation, err := tokenHandler(token)

				if err != nil {
					continue
				}

				blocksLen := len(blocks)

				switch tokenType {
				case TOKEN_DO:
					blockOperation := &BlockOperation{
						_type:     OP_BLOCK,
						block:     &Program{},
						elseBlock: &Program{},
					}

					blocks.Push(blockOperation)

				case TOKEN_ELSE:
					blocks[blocksLen-1].elseBlock.Push(operation)

				case TOKEN_END:
					block, err := blocks.Pop()

					if err != nil {
						return nil, err
					}

					program.Push(block)

				default:
					if blocksLen == 0 {
						program.Push(operation)
						break
					}

					block := blocks[blocksLen-1]

					if len(*block.elseBlock) != 0 {
						block.elseBlock.Push(operation)
					} else {
						block.block.Push(operation)
					}
				}

				found = true
				break
			}

			if !found {
				return nil, fmt.Errorf(
					"Token error in %d:%d - '%s' is not a valid token.",
					lineNumber+1, colNumber+1,
					token,
				)
			}

		}
	}

	return &program, nil
}

func NewProgramFromFile(filename string) (*Program, error) {
	rawLines, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(rawLines), "\n")

	program, err := getProgram(lines)

	if err != nil {
		log.Fatal(err)
	}

	return program, nil
}
