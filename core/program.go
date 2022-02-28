package core

import (
	"fmt"
	"os"
	"strings"
)

type Program []*Operation

func (p *Program) IsEmpty() bool {
	return len(*p) == 0
}

func (p *Program) Push(operation *Operation) {
	*p = append(*p, operation)
}

func (p *Program) Pop() (*Operation, error) {
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

func NewProgramFromFile(filename string) (*Program, error) {
	lines, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var program Program

	for _, line := range strings.Split(string(lines), "\n") {
		for _, token := range strings.Split(line, " ") {
			// TODO: check for exaustive tokens lookup (for else?)
			for _, tokenHandler := range REGISTERED_TOKENS {
				operation, err := tokenHandler(token)

				if err != nil {
					continue
				}

				program.Push(operation)
				break
			}
		}
	}

	return &program, nil
}
