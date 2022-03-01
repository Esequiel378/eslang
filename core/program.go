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

type BlockStack []BlockOperation

func (bs *BlockStack) IsEmpty() bool {
	return len(*bs) == 0
}

func (bs *BlockStack) Push(value BlockOperation) {
	*bs = append(*bs, value)
}

func (bs *BlockStack) Last() BlockOperation {
	return (*bs)[len(*bs)-1]
}

func (bs *BlockStack) Pop() (BlockOperation, error) {
	if bs.IsEmpty() {
		return nil, fmt.Errorf("can not perform `ProgramStack.Pop()`, stack is empty.")
	}

	// Get the index of the top most element.
	index := len(*bs) - 1
	// Index into the slice and obtain the element.
	value := (*bs)[index]
	// Remove it from the stack by slicing it off.
	*bs = (*bs)[:index]

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
			token = strings.Trim(token, " ")
			found := false

			for _, tokenHandler := range REGISTERED_TOKENS {
				operation, err := tokenHandler(token, &blocks)

				if err != nil {
					continue
				}

				if operation == nil {
					found = true
					break
				}

				found = true

				if blocks.IsEmpty() {
					program.Push(operation)
				} else {
					blocks.Last().PushIntoBlocks(operation)
				}

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

	// fmt.Println((*program)[1])

	if err != nil {
		log.Fatal(err)
	}

	return program, nil
}
